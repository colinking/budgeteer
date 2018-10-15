import * as React from 'react'
import { ProviderProps } from './Provider'

const defaultState: ProviderProps = {
  user: undefined
}

const { Provider, Consumer } = React.createContext(defaultState)

export { Provider, Consumer, defaultState }
