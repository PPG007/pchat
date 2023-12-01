import { createHashRouter } from "react-router-dom";
import App from "../App.tsx";
import { RegisterApplication } from "../components";

const router = createHashRouter([
  {
    path: '/',
    element: <App/>,
    children: [
      {
        element: <RegisterApplication/>,
        index: true,
      }
    ]
  }
])

export default router;