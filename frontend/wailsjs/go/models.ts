export namespace main {
	
	export class Data {
	    url: string;
	    refresh_token: string;
	    timestamp: number;
	    code: number;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.refresh_token = source["refresh_token"];
	        this.timestamp = source["timestamp"];
	        this.code = source["code"];
	        this.message = source["message"];
	    }
	}
	
	
	
	export class ResultQRLoginStatus {
	    data: Data;
	    message: string;
	    code: number;
	
	    static createFrom(source: any = {}) {
	        return new ResultQRLoginStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = this.convertValues(source["data"], Data);
	        this.message = source["message"];
	        this.code = source["code"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	

}

