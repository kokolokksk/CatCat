import { ReactElement } from 'react';
import { useLocation } from 'react-router-dom';
import DanmuWindow from '../pages/DanmuWindow';
import LivePreview from '../pages/LivePreview';
import Setting from '../pages/Setting';
import Yin from '../pages/Yin';
import Loading from '../pages/Loading';

const RouteConfig = (_props: any) => {
  const location = useLocation();
  const { search } = location;

  // 视图配置
  const viewsConfig = (pathname: string): { [key: string]: JSX.Element } => {
    return {
      notFind: <Setting />,
      dmWindow: <DanmuWindow />,
      yin: <Yin />,
      main: <Setting />,
      livePreview: <LivePreview />,
      loading: <Loading path={pathname} />,
    };
  };

  /**
   * 获取视图
   * @returns
   */
  const selectView = () => {
    const name: string = location.pathname.replace('/', '');
    if (name.includes("?loading")) {
      return viewsConfig(location.pathname).loading;
    }
    let view = viewsConfig(location.pathname)[name];
    console.log(name);
    if (view == null) {
      view = viewsConfig(location.pathname).notFind;
    }
    return view;
  };

  return <div>{selectView()}</div>;
};

export default RouteConfig;
