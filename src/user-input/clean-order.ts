// Third-Party libraries
import inquirer from 'inquirer'

// NodeJS
import process from 'node:process'

// Utils
import { writeMessage } from '@/utils/console'
import { getErrorMessage } from '@/utils/errors'

export async function addStylelintConfigCleanOrder(): Promise<boolean> {
  try {
    const { addStylelintConfigCleanOrder } = await inquirer.prompt({
      type: 'confirm',
      default: true,
      message: 'Do you wanna add stylelint-config-clean-order?:',
      name: 'addStylelintConfigCleanOrder'
    })

    return addStylelintConfigCleanOrder
  } catch {
    writeMessage({
      type: 'error',
      message: getErrorMessage('StylelintConfigCleanOrder')
    })

    process.exit(1)
  }
}
