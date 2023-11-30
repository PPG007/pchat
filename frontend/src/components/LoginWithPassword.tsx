import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Form, Input } from "antd";
import { FC } from "react";
import { useTranslation } from "react-i18next";
import { User } from "../engine";
import { LoginRequest } from "../pb/user/request.ts";
import { getMessage } from "../utils";
import { validateMessage } from "../utils/validator.ts";
import { LoginProps } from "./props.ts";

const LoginWithPassword: FC<LoginProps> = ({onSuccess, onLoading, onFailure}) => {
  const [form] = Form.useForm<LoginRequest>()
  const { t } = useTranslation();
  const onSubmit = async () => {
    const req = await form.validateFields()
    onLoading();
    try {
      const resp = await User.loginWithPassword(req)
      onSuccess(resp.data.token);
    } catch (e) {
      onFailure(getMessage(e, t('errorMessage.login')));
    }
  }
  return (
    <Form
      form={form}
      onFinish={onSubmit}
      validateMessages={validateMessage}
      colon={false}
    >
      <Form.Item
        rules={[{required: true, type: 'email'}]}
        name={'email'}
        label={t('email')}
      >
        <Input
          placeholder={t('emailPlaceholder')}
          prefix={<UserOutlined />}
        />
      </Form.Item>
      <Form.Item
        rules={[{required: true}]}
        name={'password'}
        label={t('password')}
      >
        <Input.Password
          placeholder={t('passwordPlaceholder')}
          prefix={<LockOutlined/>}
        />
      </Form.Item>
      <Form.Item wrapperCol={{span: 16, offset: 4}}>
        <Button
          type={'primary'}
          htmlType={'submit'}
          style={{width: '100%'}}
          size={'large'}
        >
          {
            t('login')
          }
        </Button>
      </Form.Item>
    </Form>
  )
}

export default LoginWithPassword;