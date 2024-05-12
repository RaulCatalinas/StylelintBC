// Third-Party libraries
import opener from 'opener'

// Constants
import { REPOSITORY } from '@/constants/github'

export const handlerOptionBuild = async () => {
  try {
    console.log("Generating Stylelint's configuration...")
  } catch (error) {
    console.error(error)
  }
}

export const handlerOptionCollaborate = async () => {
  try {
    console.log('Opening the GitHub repository...')

    setTimeout(() => opener(REPOSITORY), 500)
  } catch (error) {
    console.error(error)
  }
}
