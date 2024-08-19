/**
 * Namespace Api
 *
 * All backend api type
 */
declare namespace Api {
  namespace Common {
    /** common params of paginating */
    interface PaginatingCommonParams {
      /** current page number */
      current: number;
      /** page size */
      size: number;
      /** total count */
      total: number;
    }

    /** common params of paginating query list data */
    interface PaginatingQueryRecord<T = any> extends PaginatingCommonParams {
      records: T[];
    }

    /** common search params of table */
    type CommonSearchParams = Pick<Common.PaginatingCommonParams, 'current' | 'size'>;

    /**
     * enable status
     *
     * - "1": enabled
     * - "2": disabled
     */
    type EnableStatus = '1' | '2';

    /** common record */
    type CommonRecord<T = any> = {
      id: number;
      created_at: string;
    } & T;
  }

  namespace Form {
    interface LoginForm {
      username: string;
      password: string;
      code: string;
      captchaId: string;
    }
  }

  /**
   * namespace Auth
   *
   * backend api module: "auth"
   */
  namespace Auth {
    interface LoginToken {
      token: string;
    }

    interface Captcha {
      id: string;
      img: string;
    }

    interface UserInfo {
      userId: string;
      userName: string;
      roles: string[];
      buttons: string[];
    }
  }

  namespace Home {
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

  namespace UserManagement {
    type User = Common.CommonRecord<{
      username: string;
      password: string;
      name: string;
      email: string;
      licensed_devices: number;
      note: string;
      status: number;
      is_admin: boolean;
    }>;
    type UserList = Common.PaginatingQueryRecord<User>;
    interface Sessions {
      total: number;
      list: Session[];
    }
    interface Session {
      id: number;
      username: string;
      rustdesk_id: string;
      expired: string;
      created_at: string;
    }
  }

  /**
   * namespace Route
   *
   * backend api module: "route"
   */
  namespace Route {
    type ElegantConstRoute = import('@elegant-router/types').ElegantConstRoute;

    interface MenuRoute extends ElegantConstRoute {
      id: string;
    }

    interface UserRoute {
      routes: MenuRoute[];
      home: import('@elegant-router/types').LastLevelRouteKey;
    }
  }
}
