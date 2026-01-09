export namespace launcher {
	
	export class AccountType {
	    type: string;
	    name: string;
	    uuid: string;
	    access_token?: string;
	    refresh_token?: string;
	    client_token?: string;
	    server?: string;
	    base_code?: string;
	    head_skin: string;
	
	    static createFrom(source: any = {}) {
	        return new AccountType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.name = source["name"];
	        this.uuid = source["uuid"];
	        this.access_token = source["access_token"];
	        this.refresh_token = source["refresh_token"];
	        this.client_token = source["client_token"];
	        this.server = source["server"];
	        this.base_code = source["base_code"];
	        this.head_skin = source["head_skin"];
	    }
	}
	export class AccountList {
	    accounts: AccountType[];
	
	    static createFrom(source: any = {}) {
	        return new AccountList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.accounts = this.convertValues(source["accounts"], AccountType);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class ExceptionHandler_NovaPlus_module_launcher_AccountList_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: AccountList;
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler_NovaPlus_module_launcher_AccountList_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], AccountList);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ExceptionHandler_NovaPlus_module_launcher_AccountType_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: AccountType;
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler_NovaPlus_module_launcher_AccountType_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], AccountType);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class JavaConfig {
	    path: string;
	    version: string;
	    arch: string;
	    vendor: string;
	
	    static createFrom(source: any = {}) {
	        return new JavaConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.version = source["version"];
	        this.arch = source["arch"];
	        this.vendor = source["vendor"];
	    }
	}
	export class ExceptionHandler_NovaPlus_module_launcher_JavaConfig_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: JavaConfig;
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler_NovaPlus_module_launcher_JavaConfig_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], JavaConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class JavaConfigs {
	    java: JavaConfig[];
	
	    static createFrom(source: any = {}) {
	        return new JavaConfigs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.java = this.convertValues(source["java"], JavaConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ExceptionHandler_NovaPlus_module_launcher_JavaConfigs_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: JavaConfigs;
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler_NovaPlus_module_launcher_JavaConfigs_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], JavaConfigs);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MCConfig {
	    path: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new MCConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.name = source["name"];
	    }
	}
	export class MCConfigs {
	    mc: MCConfig[];
	
	    static createFrom(source: any = {}) {
	        return new MCConfigs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.mc = this.convertValues(source["mc"], MCConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ExceptionHandler_NovaPlus_module_launcher_MCConfigs_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: MCConfigs;
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler_NovaPlus_module_launcher_MCConfigs_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], MCConfigs);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ExceptionHandler___NovaPlus_module_launcher_AccountType_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: AccountType[];
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler___NovaPlus_module_launcher_AccountType_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], AccountType);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class HomePageStruct {
	    name: string;
	    path: string;
	
	    static createFrom(source: any = {}) {
	        return new HomePageStruct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	    }
	}
	export class ExceptionHandler___NovaPlus_module_launcher_HomePageStruct_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: HomePageStruct[];
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler___NovaPlus_module_launcher_HomePageStruct_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], HomePageStruct);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class IPv6Struct {
	    ip: string;
	    success: boolean;
	
	    static createFrom(source: any = {}) {
	        return new IPv6Struct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ip = source["ip"];
	        this.success = source["success"];
	    }
	}
	export class ExceptionHandler___NovaPlus_module_launcher_IPv6Struct_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: IPv6Struct[];
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler___NovaPlus_module_launcher_IPv6Struct_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = this.convertValues(source["data"], IPv6Struct);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ExceptionHandler___string_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: string[];
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler___string_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}
	export class ExceptionHandler_interface____ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler_interface____(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}
	export class ExceptionHandler_string_ {
	    code: number;
	    status: boolean;
	    message: string;
	    data: string;
	
	    static createFrom(source: any = {}) {
	        return new ExceptionHandler_string_(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.status = source["status"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}
	
	
	
	
	

}

