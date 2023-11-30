import { Button, Form, Input, Space } from "antd";
import { FC, Fragment, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { validateMessage } from "../utils/validator.ts";
import { LoginProps } from "./props.ts";

const LoginWithCaptcha: FC<LoginProps> = () => {
  const {t} = useTranslation();
  const [form] = Form.useForm();
  return (
    <Fragment>
      <Form
        form={form}
        colon={false}
        labelCol={{
         span: 4
        }}
        wrapperCol={{
          span: 20
        }}
        validateMessages={validateMessage}
      >
        <Form.Item
          name={'email'}
          label={t('email')}
          rules={[{required: true, type: 'email'}]}
        >
          <Input/>
        </Form.Item>
        <Form.Item
          name={'captcha'}
          label={t('captcha')}
          rules={[{required: true}]}
        >
          <Space.Compact block>
            <Input/>
            <CaptchaButton/>
          </Space.Compact>
        </Form.Item>
        <Form.Item wrapperCol={{span:16, offset: 4}}>
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
    </Fragment>
  )
}

const CaptchaButton: FC = () => {
  const [second, setSecond] = useState(0);
  const {t} = useTranslation();
  const interval = useRef<NodeJS.Timeout>();
  const onClick = () => {
    setSecond(60);
    interval.current = setInterval(() => {
      setSecond((s) => {
        if (s === 1) {
          clearInterval(interval.current);
          return 0
        }
        return s-1;
      });
    }, 1000);
  }
  return (
    <Button
      disabled={second !== 0}
      onClick={onClick}
    >
      {
        second === 0 ? t('getCaptcha') : second
      }
    </Button>
  )
}

export default LoginWithCaptcha;