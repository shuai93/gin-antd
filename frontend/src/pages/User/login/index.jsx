import {LockOutlined, UserOutlined,} from '@ant-design/icons';
import {Alert, Tabs} from 'antd';
import React, {useState} from 'react';
import ProForm, {ProFormCheckbox, ProFormText} from '@ant-design/pro-form';
import {connect} from 'umi';
import styles from './index.less';

const LoginMessage = ({content}) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const Login = (props) => {
  const {userLogin = {}, submitting} = props;
  const {status, type: loginType} = userLogin;
  const [type, setType] = useState('account');

  const handleSubmit = (values) => {
    const {dispatch} = props;
    dispatch({
      type: 'login/login',
      payload: {...values, type},
    });
  };

  return (
    <div className={styles.main}>
      <ProForm
        initialValues={{
          autoLogin: true,
        }}
        submitter={{
          render: (_, dom) => dom.pop(),
          submitButtonProps: {
            loading: submitting,
            size: 'large',
            style: {
              width: '100%',
            },
          },
        }}
        onFinish={(values) => {
          handleSubmit(values);
          return Promise.resolve();
        }}
      >
        <Tabs activeKey={type} onChange={setType}>
          <Tabs.TabPane
            key="account"
            tab='账户密码登录'
          />
          {/*<Tabs.TabPane*/}
          {/*  key="mobile"*/}
          {/*  tab={intl.formatMessage({*/}
          {/*    id: 'pages.login.phoneLogin.tab',*/}
          {/*    defaultMessage: '手机号登录',*/}
          {/*  })}*/}
          {/*/>*/}
        </Tabs>

        {status === 'error' && loginType === 'account' && !submitting && (
          <LoginMessage
            content='账户或密码错误'
          />
        )}
        {type === 'account' && (
          <>
            <ProFormText
              name="userName"
              fieldProps={{
                size: 'large',
                prefix: <UserOutlined className={styles.prefixIcon}/>,
              }}
              placeholder='用户名'
              rules={[
                {
                  required: true,
                  message: '用户名是必填项！',
                },
              ]}
            />
            <ProFormText.Password
              name="password"
              fieldProps={{
                size: 'large',
                prefix: <LockOutlined className={styles.prefixIcon}/>,
              }}
              placeholder='密码'
              rules={[
                {
                  required: true,
                  message: '密码是必填项！',
                },
              ]}
            />
          </>
        )}

        <div
          style={{
            marginBottom: 24,
          }}
        >
          <ProFormCheckbox noStyle name="autoLogin">
            自动登录
          </ProFormCheckbox>
          <a
            style={{
              float: 'right',
            }}
          >
            忘记密码 ?
          </a>
        </div>
      </ProForm>
      {/*<Space className={styles.other}>*/}
      {/*  其他登录方式 :*/}
      {/*  <AlipayCircleOutlined className={styles.icon} />*/}
      {/*  <TaobaoCircleOutlined className={styles.icon} />*/}
      {/*  <WeiboCircleOutlined className={styles.icon} />*/}
      {/*</Space>*/}
    </div>
  );
};

export default connect(({login, loading}) => ({
  userLogin: login,
  submitting: loading.effects['login/login'],
}))(Login);
