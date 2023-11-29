import { Card } from "antd";
import { FC } from "react";
import { LoginWithPassword } from "../components";
import styles from './styles/login.module.less';
const Login: FC = () => {
  return (
    <div>
      <Card className={styles.loginCard}>
        <LoginWithPassword/>
      </Card>
    </div>
  )
}

export default Login;