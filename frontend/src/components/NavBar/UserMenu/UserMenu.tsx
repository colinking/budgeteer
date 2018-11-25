import { Button, Pane, Popover, Text } from 'evergreen-ui'
import * as React from 'react'
import { User } from '../../../lib/auth'
// import AvatarDropdown from './AvatarDropdown'
// import Menu from './Menu'

declare interface UserMenuProps {
  user: User
}

export default class UserMenu extends React.Component<UserMenuProps> {
  public render() {
    // const props = this.props

    return (
      // <Popover content={<Menu user={props.user}/>}>
      //   <AvatarDropdown user={props.user} />
      // </Popover>
      <Popover
        content={
          <Pane
            width={240}
            height={240}
            display="flex"
            alignItems="center"
            justifyContent="center"
            flexDirection="column"
          >
            <Text>PopoverContent</Text>
            {/* <Menu user={props.user}/> */}
          </Pane>
        }
      >
        <Button>Trigger Popover</Button>
      </Popover>
    )
  }
}
