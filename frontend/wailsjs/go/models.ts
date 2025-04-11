export namespace main {
	
	export class configData {
	    Ip: any;
	    Path: any;
	    CodeReadonly: any;
	    CodeReadwrite: any;
	    LockUploaddir: any;
	
	    static createFrom(source: any = {}) {
	        return new configData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Ip = source["Ip"];
	        this.Path = source["Path"];
	        this.CodeReadonly = source["CodeReadonly"];
	        this.CodeReadwrite = source["CodeReadwrite"];
	        this.LockUploaddir = source["LockUploaddir"];
	    }
	}

}

