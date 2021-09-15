import React from 'react';
import {connect} from 'umi';
import Avatar from './AvatarDropdown';
import {GlobalNotice} from './NoticeIconView'
import styles from './index.less';

const ENVTagColor = {
  dev: 'orange',
  test: 'green',
  pre: '#87d068',
};

const GlobalHeaderRight = (props) => {
  const {theme, layout} = props;
  let className = styles.right;

  if (theme === 'dark' && layout === 'top') {
    className = `${styles.right}  ${styles.dark}`;
  }

  return (
    <div className={className}>
      <GlobalNotice/>
      <Avatar/>
    </div>
  );
};

export default connect(({settings}) => ({
  theme: settings.navTheme,
  layout: settings.layout,
}))(GlobalHeaderRight);
