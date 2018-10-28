import { Text } from 'evergreen-ui'
import * as React from 'react'

export interface WithError {
  error: string,
  description: string
}

export default class Error extends React.Component<WithError> {
  public render() {
    return (
      <Text>
        Error! {this.props.error} -- {this.props.description}
      </Text>
    )
  }
}
