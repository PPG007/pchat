import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from "react-redux";
import { RouterProvider } from "react-router-dom";
import {App} from 'antd';
import '../i18n';
import router from "./routes";
import store from "./store";

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App>
      <Provider store={store}>
        <RouterProvider router={router}/>
      </Provider>
    </App>
  </StrictMode>,
)
