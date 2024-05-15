// Third-Party libraries
import opener from 'opener'

// Constants
import { REPOSITORY } from '@/constants/github'

// Utils
import { writeMessage } from '@/utils/console'
import { getErrorMessage } from '@/utils/errors'
import { generateStylelintConfig } from '@/utils/stylelint'
import { exists } from '@/utils/user-os'

// User-Input
import { addStylelintConfigCleanOrder } from '@/user-input/clean-order'
import { getPackageManger } from '@/user-input/package-managers'
import { usingVSCodeEditor } from '@/user-input/using-vscode'

export const handlerOptionBuild = async () => {
  try {
    const packageJsonPath = `${process.cwd()}/package.json`

    const existPackageJsonInTheCurrentDirectory = await exists(packageJsonPath)

    if (!existPackageJsonInTheCurrentDirectory) {
      writeMessage({
        type: 'error',
        message: getErrorMessage('NotFound')
      })

      process.exit(1)
    }

    const packageManagerToUse = await getPackageManger()
    const usingVSCEditor = await usingVSCodeEditor()
    const useStylelintConfigCleanOrder = await addStylelintConfigCleanOrder()

    await generateStylelintConfig({
      packageManagerToUse,
      usingVSCodeEditor: usingVSCEditor,
      useStylelintConfigCleanOrder
    })
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('Default')
    })

    process.exit(1)
  }
}

export const handlerOptionCollaborate = async () => {
  console.log('Opening the GitHub repository...')

  setTimeout(() => opener(REPOSITORY), 500)
}
