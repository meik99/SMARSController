import {Vue} from "vue-class-component";

export class Authentication {
    hasAuthCode(): boolean {
        const code = localStorage.getItem("code");
        return !!code;
    }

    getAuthCode(): string {
        const code = localStorage.getItem("code");
        return code ? code : '';
    }

    setAuthCode(code: string) {
        localStorage.setItem('code', code);
    }

    requestAuthCode(): void {
        localStorage.removeItem('code');
        let pathname = `${location.protocol}//${location.host}${location.pathname}`
        if (pathname.endsWith('/')) {
            pathname = pathname.substring(0, pathname.length - 1);
        }
        console.log(pathname)
        location.href = process.env.VUE_APP_COFFEE_AUTH + "?state=" +
            encodeURI(pathname);
    }

    checkAuthStatusForApp(app: Vue): void {
        if(!app.$route.query.code && !localStorage.getItem("code")) {
            this.requestAuthCode();
        } else if(!localStorage.getItem("code")) {
            const codeOrNull = app.$route.query.code;
            const code = (codeOrNull as string) ?
                (codeOrNull as string) :
                '';
            localStorage.setItem("code", code);
            location.href = process.env.VUE_APP_COFFEE_ALARM_UI;
        }
    }

    clear() {
        localStorage.removeItem('code');
    }
}