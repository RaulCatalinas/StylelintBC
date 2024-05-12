import { stylelintbc } from '@/constants/stylelintbc'
import { VERSION } from '@/constants/version'

export function configureCLI() {
  stylelintbc
    .description('Command line for easy Stylelint configuration')
    .version(VERSION)
    .showHelpAfterError()
}
