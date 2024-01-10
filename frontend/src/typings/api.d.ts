// 后端接口返回的数据类型

/** 后端返回的用户权益相关类型 */
declare namespace ApiAuth {
  /** 返回的token和刷新token */
  interface Token {
    token: string;
  }
  /** 返回验证码 */
  interface Captcha {
    id: string;
    img: string;
  }
  /** 返回的用户信息 */
  type UserInfo = Auth.UserInfo;
}

declare namespace ApiDashboard {
  interface Stat {
    userCount: number;
    peerCount: number;
    onlineCount: number;
    visitsCount: number;
  }

  interface LineChart {
    xAxis: string[];
    users: number[];
    peer: number[];
  }

  interface PieChart {
    name: string;
    value: number;
  }
}
