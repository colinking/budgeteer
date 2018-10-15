import 'ianstormtaylor-reset'
import 'normalize.css/normalize.css'
import 'react-virtualized/styles.css'

import * as React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom'

import App from './components/App'
import './index.scss'
import registerServiceWorker from './registerServiceWorker'

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root') as HTMLElement
)
registerServiceWorker()
