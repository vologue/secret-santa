import React, { Component, useState } from 'react';
import './styles/loading.styles.scss';
import {Spin, Affix} from 'antd';
import { LoadingOutlined } from '@ant-design/icons';

export var LoadingScreen = (status) =>{
  const antIcon = <LoadingOutlined style={{ fontSize: 24 }} spin />;
  if (status.loading){
    return (
      <div className="loading-container" id="overlay">
        <Spin indicator={antIcon} tip={status.loadingText}/>
      </div>
    )
  }
  else {
    return null
  }
}
