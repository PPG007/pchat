import { Col, Progress, QRCode, Row } from "antd";
import { FC, Fragment, useEffect, useRef, useState } from "react";
import { LoginProps } from "./props.ts";
import {v4 as uuid} from "uuid";

const LoginWithQRCode: FC<LoginProps> = () => {
  const [percent, setPercent] = useState(100);
  const [value, setValue] = useState(uuid());
  const intervalId = useRef<NodeJS.Timeout>();
  useEffect(() => {
    const id = setInterval(() => {
      setPercent((p) => {
        if (p <= 1.7) {
          fetchCodeValue();
          return 0;
        }
        return p - 1.7;
      })
    }, 1000);
    intervalId.current = id;
    return () => {clearInterval(id)}
  }, [value])
  const fetchCodeValue = () => {
    setPercent(100);
    clearInterval(intervalId.current);
    setValue(uuid());
  }
  // TODO:
  return (
    <Fragment>
      <Row>
        <Col span={16} offset={8}>
          <QRCode value={value}/>
        </Col>
      </Row>
      <Row>
        <Col span={20} offset={2}>
          <Progress
            percent={percent}
            showInfo={false}
          />
        </Col>
      </Row>
    </Fragment>
  )
}

export default LoginWithQRCode;