import { majorScale, Spinner } from 'evergreen-ui'
import * as React from 'react'

const Loader = () => {
  return (
    <Spinner marginX="auto" marginY={majorScale(2)} size={majorScale(3)} />
  )
}

export default Loader
