import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Form, Input } from "antd";
import { FC } from "react";
import { useTranslation } from "react-i18next";
import { LoginRequest } from "../pb/user/request.ts";
import { validateMessage } from "../utils/validator.ts";

const LoginWithPassword: FC = () => {
  const [form] = Form.useForm<LoginRequest>()
  const { t } = useTranslation();
  const onSubmit = async () => {
    const result = await form.validateFields()
    console.log(result);
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
      <Form.Item>
        <Button type={'primary'} htmlType={'submit'} style={{width: '100%'}}>
          {
            t('login')
          }
        </Button>
      </Form.Item>
    </Form>
  )
}

export default LoginWithPassword;