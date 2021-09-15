import React from 'react';
import {PageContainer} from '@ant-design/pro-layout';
import {Card, Typography} from 'antd';
// noinspection NpmUsedModulesInstalled
import DPlayer from "react-dplayer";


export default () => {
  const dPlayerOptions = {
    video: {url: 'http://static.smartisanos.cn/common/video/t1-ui.mp4'}
  }
  return (
    <PageContainer>
      <Card
        title="DPlayer"
      >
        <Typography.Text strong>

          <DPlayer
            options={dPlayerOptions}
          />
        </Typography.Text>

      </Card>


    </PageContainer>
  );
};
