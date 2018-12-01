import { Popover, Avatar, Pane, majorScale, minorScale, Icon } from 'evergreen-ui'
import * as React from 'react'
import { User } from '../../../lib/auth'
import Menu from './Menu'

declare interface AvatarMenuProps {
  user: User
}

export default class AvatarMenu extends React.Component<AvatarMenuProps> {
  public render() {
    const props = this.props
    const avatarProps = {
      name: props.user.fullName,
      src: props.user.picture
    }

    return (
      <Popover bringFocusInside={false} content={<Menu user={props.user}/>}>
        <Pane
          tabIndex={0}
          role="button"
          display="flex"
          cursor="pointer"
          alignItems="center"
        >
          <Avatar size={majorScale(4)} color="green" isSolid {...avatarProps} />
          <Icon
            icon="caret-down"
            marginLeft={minorScale(1)}
            size={minorScale(4)}
            color="inherit"
          />
        </Pane>
      </Popover>
    )
  }
}
