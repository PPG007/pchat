import { App, Card, Spin } from "antd";
import { CardTabListType } from "antd/es/card";
import { FC, useState } from "react";
import { useTranslation } from "react-i18next";
import { useDispatch } from "react-redux";
import { LoginWithCaptcha, LoginWithPassword, LoginWithQRCode } from "../components";
import { LoginProps } from "../components/props.ts";
import { Dispatch, login } from "../store";
import styles from './styles/login.module.less';
const Login: FC = () => {
  const dispatch = useDispatch<Dispatch>();
  const {t} = useTranslation();
  const [tab, setTab] = useState('password');
  const [isLoading, setIsLoading] = useState(false);
  const {message} = App.useApp();
  const onLogin = (token: string) => {
    dispatch(login({token}));
    setIsLoading(false);
  }
  const onFailure = (m: string) => {
    setIsLoading(false);
    message.error(m);
  }
  const onLoading = () => {
    setIsLoading(true);
  }
  const tabs: Array<CardTabListType> = [
    {
      key: 'password',
      label: t('login', {context: 'password'}),
      disabled: isLoading,
    },
    {
      key: 'qrcode',
      label: t('login', {context: 'qrcode'}),
      disabled: isLoading,
    },
    {
      key: 'captcha',
      label: t('login', {context: 'captcha'}),
      disabled: isLoading,
    }
  ];

  return (
    <div>
      <Card
        className={styles.loginCard}
        tabList={tabs}
        onTabChange={(key) => {setTab(key)}}
        activeTabKey={tab}
        hoverable
        tabProps={{
          centered: true,
          destroyInactiveTabPane: true,
        }}
      >
        <Spin spinning={isLoading}>
          <LoginFC props={{
            onSuccess: onLogin,
            onFailure,
            onLoading,
          }} tab={tab}/>
        </Spin>
      </Card>
    </div>
  )
}

const LoginFC: FC<{props: LoginProps, tab: string}> = ({props, tab}) => {
  switch (tab) {
    case 'qrcode':
      return <LoginWithQRCode {...props}/>
    case 'captcha':
      return <LoginWithCaptcha {...props}/>
    default:
      return <LoginWithPassword {...props}/>
  }
}

export default Login;