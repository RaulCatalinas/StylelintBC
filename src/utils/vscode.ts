// NodeJS
import fs from 'node:fs/promises'
import process from 'node:process'

// Utils
import { writeMessage } from './console'
import { getErrorMessage } from './errors'
import { createEmptyJsonFile, exists } from './user-os'

// Constants
import { UTF8_ENCODING } from '@/constants/encoding'

// Types
import type { Extensions, Settings } from '@/types/vscode'

export async function configureVSCode() {
  try {
    writeMessage({
      type: 'info',
      message: 'Configuring VSCode...'
    })

    const existVSCodeSettingsFile = await exists(
      `${process.cwd()}/.vscode/settings.json`
    )

    if (!existVSCodeSettingsFile) {
      await createEmptyJsonFile('.vscode/settings.json')
    }

    const vscodeSettings = await fs.readFile('.vscode/settings.json', {
      encoding: UTF8_ENCODING
    })

    const vscodeSettingsObject: Settings = JSON.parse(vscodeSettings)

    vscodeSettingsObject['[css]'] = JSON.stringify({
      'editor.defaultFormatter': 'stylelint.vscode-stylelint'
    })

    await fs.writeFile(
      '.vscode/settings.json',
      JSON.stringify(vscodeSettingsObject, null, 2),
      {
        encoding: UTF8_ENCODING
      }
    )

    writeMessage({
      type: 'info',
      message: 'VSCode configured successfully'
    })
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('VSCodeConfig')
    })

    process.exit(1)
  }
}

export async function addRecommendedExtension(extension: string) {
  try {
    writeMessage({
      type: 'info',
      message: 'Adding recommended extensions...'
    })

    const existVSCodeSettingsFile = await exists(
      `${process.cwd()}/.vscode/extensions.json`
    )

    if (!existVSCodeSettingsFile) {
      await createEmptyJsonFile('.vscode/extensions.json')
    }

    const vscodeExtensions = await fs.readFile('.vscode/extensions.json', {
      encoding: UTF8_ENCODING
    })

    const vscodeExtensionsObject: Extensions = JSON.parse(vscodeExtensions)

    if (vscodeExtensionsObject.recommendations === undefined) {
      vscodeExtensionsObject.recommendations = []

      await fs.writeFile(
        '.vscode/extensions.json',
        JSON.stringify(vscodeExtensionsObject, null, 2),
        {
          encoding: UTF8_ENCODING
        }
      )
    }

    vscodeExtensionsObject.recommendations.push(extension)

    await fs.writeFile(
      '.vscode/extensions.json',
      JSON.stringify(vscodeExtensionsObject, null, 2),
      {
        encoding: UTF8_ENCODING
      }
    )

    writeMessage({
      type: 'info',
      message: 'Recommended extensions added successfully'
    })
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('AddRecommendedExtensions')
    })

    process.exit(1)
  }
}
