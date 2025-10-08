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

    /** common record */
    type CommonRecord<T = any> = {
      id?: number;
      created_at?: string;
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
      deviceCount: number;
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
      login_verify: string;
      tfa_secret: string;
      note: string;
      status: number;
      is_admin: boolean;
      admin_status: number;
      tfa_code: string;
    }>;
    type UserList = Common.PaginatingQueryRecord<User>;

    type UserSearchParams = CommonType.RecordNullable<
      Pick<Api.UserManagement.User, 'username' | 'name' | 'email' | 'status' | 'admin_status' | 'created_at'> &
        Api.Common.CommonSearchParams
    >;

    type Session = Common.CommonRecord<{
      username: string;
      rustdesk_id: string;
      expired: string;
    }>;
    type SessionList = Common.PaginatingQueryRecord<Session>;

    type SessionSearchParams = CommonType.RecordNullable<
      Pick<Api.UserManagement.Session, 'username' | 'created_at'> & Api.Common.CommonSearchParams
    >;
  }

  namespace Devices {
    type Device = Common.CommonRecord<{
      rustdesk_id: string;
      hostname: string;
      username: string;
      uuid: string;
      version:string;
      os:string;
      memory:string;
      created_at: string;
    }>;
    type DevicesList = Common.PaginatingQueryRecord<Device>;
    type DeviceSearchParams = CommonType.RecordNullable<
      Pick<
        Api.Devices.Device,
        'username' | 'hostname' | 'rustdesk_id'
      > &
        Api.Common.CommonSearchParams
    >;
  }

  namespace Audit {
    type AuditLog = Common.CommonRecord<{
      username: string;
      conn_id: string;
      rustdesk_id: string;
      ip: string;
      session_id: string;
      uuid: string;
      type: number;
      closed_at?: string;
    }>;
    type AuditLogList = Common.PaginatingQueryRecord<AuditLog>;
    type AuditLogSearchParams = CommonType.RecordNullable<
      Pick<
        Api.Audit.AuditLog,
        'username' | 'conn_id' | 'rustdesk_id' | 'ip' | 'session_id' | 'uuid' | 'type' | 'created_at' | 'closed_at'
      > &
        Api.Common.CommonSearchParams
    >;

    type AuditFileTransferLog = Common.CommonRecord<{
      rustdesk_id: string;
      peer_id: string;
      path: string;
      uuid: string;
      type: number;
    }>;
    type AuditFileTransferList = Common.PaginatingQueryRecord<AuditFileTransferLog>;
    type AuditFileTransferLogSearchParams = CommonType.RecordNullable<
      Pick<Api.Audit.AuditFileTransferLog, 'rustdesk_id' | 'peer_id' | 'uuid' | 'type' | 'created_at'> &
        Api.Common.CommonSearchParams
    >;
  }

  namespace System {
    type MailTemplate = Common.CommonRecord<{
      name: string;
      type: number;
      subject: string;
      contents: string;
    }>;
    type MailTemplateList = Common.PaginatingQueryRecord<MailTemplate>;
    type MailTemplateSearchParams = CommonType.RecordNullable<
      Pick<Api.System.MailTemplate, 'name' | 'type' | 'subject' | 'contents' | 'created_at'> &
        Api.Common.CommonSearchParams
    >;

    type MailLog = Common.CommonRecord<{
      username: string;
      uuid: string;
      subject: string;
      from: string;
      to: string;
      status: number;
    }>;
    type MailLogList = Common.PaginatingQueryRecord<MailLog>;
    type MailLogSearchParams = CommonType.RecordNullable<
      Pick<Api.System.MailLog, 'username' | 'uuid' | 'subject' | 'from' | 'to' | 'status' | 'created_at'> &
        Api.Common.CommonSearchParams
    >;
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
