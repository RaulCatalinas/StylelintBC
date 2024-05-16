// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Utils
import { writeMessage } from './console'
import { installDependencies } from './dependencies'
import { getErrorMessage } from './errors'

// Types
import type { PackageManager } from '@/types/package-mangers'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'
import { addRecommendedExtension, configureVSCode } from './vscode'

interface Props {
  packageManagerToUse: PackageManager
  usingVSCodeEditor: boolean
  useStylelintConfigCleanOrder: boolean
}

export async function generateStylelintConfig({
  packageManagerToUse,
  usingVSCodeEditor,
  useStylelintConfigCleanOrder
}: Props) {
  try {
    writeMessage({
      type: 'config',
      message: "Generating Stylelint's configuration..."
    })

    await installDependencies({
      packageManagerToUse,
      packagesToInstall: [
        'stylelint',
        'stylelint-config-standard',
        `${useStylelintConfigCleanOrder ? 'stylelint-config-clean-order' : ''}`
      ].filter(packageToInstall => packageToInstall !== '')
    })

    writeMessage({
      type: 'info',
      message: 'Creating configuration file...'
    })

    await fs.writeFile(
      '.stylelintrc.json',
      JSON.stringify({
        extends: [
          'stylelint-config-standard',
          `${useStylelintConfigCleanOrder ? 'stylelint-config-clean-order' : ''}`
        ].filter(string => string !== '')
      }),
      {
        encoding: UTF8_ENCODING
      }
    )

    writeMessage({
      type: 'info',
      message: 'Configuration file (.stylelintrc.json) created successfully'
    })

    if (usingVSCodeEditor) {
      await Promise.all([
        configureVSCode(),
        addRecommendedExtension('stylelint.vscode-stylelint')
      ])
    }

    writeMessage({
      type: 'success',
      message: "Stylelint's configuration generated successfully"
    })
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('Stylelint')
    })

    process.exit(1)
  }
}
