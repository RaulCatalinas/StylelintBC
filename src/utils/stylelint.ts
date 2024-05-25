// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Utils
import { writeMessage } from './console'
import { installDependencies } from './dependencies'
import { getErrorMessage } from './errors'
import { addRecommendedExtension, configureVSCode } from './vscode'

// Types
import type { PackageManager } from '@/types/package-mangers'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'

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

    await Promise.all(
      [
        installDependencies({
          packageManagerToUse,
          packagesToInstall: [
            'stylelint',
            'stylelint-config-standard',
            `${useStylelintConfigCleanOrder ? 'stylelint-config-clean-order' : ''}`
          ].filter(packageToInstall => packageToInstall !== '')
        }),

        createStylelintConfigFiles(useStylelintConfigCleanOrder),

        usingVSCodeEditor ? configureVSCode() : null,

        usingVSCodeEditor
          ? addRecommendedExtension('stylelint.vscode-stylelint')
          : null
      ].filter(promise => promise != null)
    )

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

async function createStylelintConfigFiles(
  useStylelintConfigCleanOrder: boolean
) {
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
    type: 'success',
    message: 'Configuration file (.stylelintrc.json) created successfully'
  })
}
