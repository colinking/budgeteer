import { Avatar, Icon, majorScale, minorScale, Pane } from 'evergreen-ui'
import * as React from 'react'
import { User } from '../../../../lib/auth'

declare interface AvatarProps {
  user: User
}

export default class AvatarDropdown extends React.Component<AvatarProps> {
  public render() {
    const props = this.props
    const avatarProps = {
      name: `${props.user.firstName} ${props.user.lastName}`.trim(),
      src: props.user.picture
    }

    return (
      <Pane
        tabIndex={0}
        role="button"
        display="flex"
        cursor="pointer"
        alignItems="center"
      >
        <Avatar size={majorScale(4)} color="green" {...avatarProps} />
        <Icon
          icon="caret-down"
          marginLeft={minorScale(1)}
          size={minorScale(4)}
          color="inherit"
        />
      </Pane>
    )
  }
}
