// Constants
import { stylelintbc } from '@/constants/stylelintbc'

// Controllers
import {
  handlerOptionBuild,
  handlerOptionCollaborate
} from '@/controllers/handlers-options'

export function configureOptions() {
  stylelintbc
    .option(
      '-co, --collaborate',
      'Open GitHub repository for collaboration',
      handlerOptionCollaborate
    )
    .option(
      '-b, --build',
      "Start Stylelint's configuration",
      handlerOptionBuild
    )
}

export function configureDefaultOption() {
  const options = stylelintbc.opts()

  // eslint-disable-next-line @typescript-eslint/strict-boolean-expressions
  if (!options.build && !options.collaborate && !options.b && !options.co) {
    stylelintbc.help()
  }
}
