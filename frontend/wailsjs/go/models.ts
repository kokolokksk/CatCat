export namespace main {
	
	
	
	
	export class ResultQRLoginStatus {
	    status: boolean;
	    data: any;
	    message: string;
	    code: number;
	    ts: number;
	
	    static createFrom(source: any = {}) {
	        return new ResultQRLoginStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.data = source["data"];
	        this.message = source["message"];
	        this.code = source["code"];
	        this.ts = source["ts"];
	    }
	}
	

}

