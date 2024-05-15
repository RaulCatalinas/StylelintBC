// Third-Party libraries
import inquirer from 'inquirer'

// NodeJS
import process from 'node:process'

// Utils
import { writeMessage } from '@/utils/console'
import { getErrorMessage } from '@/utils/errors'

export async function usingVSCEditor(): Promise<boolean> {
  try {
    const { usingVSCodeEditor } = await inquirer.prompt({
      type: 'confirm',
      default: true,
      message: 'Will you use VS Code as a code editor?:',
      name: 'usingVSCodeEditor'
    })

    return usingVSCodeEditor
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('VSCodeEditor')
    })

    process.exit(1)
  }
}
