import { createHashRouter } from "react-router-dom";
import App from "../App.tsx";
import { Login } from "../pages";

const router = createHashRouter([
  {
    path: '/',
    element: <App/>,
    children: [
      {
        element: <Login/>,
        index: true,
      }
    ]
  }
])

export default router;