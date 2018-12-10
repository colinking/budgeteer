import * as React from 'react'
import { UserProviderProps } from './Provider'

const defaultState: UserProviderProps = {
  user: undefined,
  refetchUser: () => null
}

const { Provider, Consumer } = React.createContext(defaultState)

export { Provider, Consumer, defaultState }
