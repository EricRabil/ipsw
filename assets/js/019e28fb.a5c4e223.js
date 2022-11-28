"use strict";(self.webpackChunkdocumentation=self.webpackChunkdocumentation||[]).push([[3526],{3905:(e,r,t)=>{t.d(r,{Zo:()=>p,kt:()=>m});var n=t(7294);function a(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function i(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function o(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?i(Object(t),!0).forEach((function(r){a(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function s(e,r){if(null==e)return{};var t,n,a=function(e,r){if(null==e)return{};var t,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||(a[t]=e[t]);return a}(e,r);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var c=n.createContext({}),l=function(e){var r=n.useContext(c),t=r;return e&&(t="function"==typeof e?e(r):o(o({},r),e)),t},p=function(e){var r=l(e.components);return n.createElement(c.Provider,{value:r},e.children)},d={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},u=n.forwardRef((function(e,r){var t=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),u=l(t),m=a,f=u["".concat(c,".").concat(m)]||u[m]||d[m]||i;return t?n.createElement(f,o(o({ref:r},p),{},{components:t})):n.createElement(f,o({ref:r},p))}));function m(e,r){var t=arguments,a=r&&r.mdxType;if("string"==typeof e||a){var i=t.length,o=new Array(i);o[0]=u;var s={};for(var c in r)hasOwnProperty.call(r,c)&&(s[c]=r[c]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var l=2;l<i;l++)o[l]=t[l];return n.createElement.apply(null,o)}return n.createElement.apply(null,t)}u.displayName="MDXCreateElement"},4690:(e,r,t)=>{t.r(r),t.d(r,{assets:()=>c,contentTitle:()=>o,default:()=>d,frontMatter:()=>i,metadata:()=>s,toc:()=>l});var n=t(7462),a=(t(7294),t(3905));const i={description:"How to extract the files you need from OTAs."},o="Parse OTAs",s={unversionedId:"guides/ota",id:"guides/ota",title:"Parse OTAs",description:"How to extract the files you need from OTAs.",source:"@site/docs/guides/ota.md",sourceDirName:"guides",slug:"/guides/ota",permalink:"/ipsw/docs/guides/ota",draft:!1,editUrl:"https://github.com/blacktop/ipsw/tree/master/www/docs/guides/ota.md",tags:[],version:"current",frontMatter:{description:"How to extract the files you need from OTAs."},sidebar:"docs",previous:{title:"Parse dyld_shared_cache",permalink:"/ipsw/docs/guides/dyld"},next:{title:"Lookup DSC Symbols",permalink:"/ipsw/docs/guides/dump_dsc_syms"}},c={},l=[{value:"Show OTA Info",id:"show-ota-info",level:4},{value:"List files in OTA",id:"list-files-in-ota",level:4},{value:"Extract file(s) from OTA payloads",id:"extract-files-from-ota-payloads",level:4}],p={toc:l};function d(e){let{components:r,...t}=e;return(0,a.kt)("wrapper",(0,n.Z)({},p,t,{components:r,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"parse-otas"},"Parse OTAs"),(0,a.kt)("h4",{id:"show-ota-info"},"Show OTA Info"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"\u276f ipsw ota --info OTA.zip\n\n[OTA Info]\n==========\nVersion        = 13.5\nBuildVersion   = 17F5054h\nOS Type        = Beta\n\nDevices\n-------\n\niPhone SE (2nd generation))\n - iPhone12,8_D79AP_17F5054h\n   - KernelCache: kernelcache.release.iphone12c\n")),(0,a.kt)("h4",{id:"list-files-in-ota"},"List files in OTA"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"\u276f ipsw ota OTA.zip | head\n   \u2022 Listing files...\n-rw-r--r-- 2020-02-15T02:24:26-05:00 0 B    .Trashes\n---------- 2020-02-15T02:20:25-05:00 0 B    .file\n-rwxr-xr-x 2020-02-15T02:23:53-05:00 0 B    etc\n-rwxr-xr-x 2020-02-15T02:24:07-05:00 0 B    tmp\n-rwxr-xr-x 2020-02-15T02:24:11-05:00 0 B    var\n-rwxrwxr-x 2020-02-15T02:20:25-05:00 109 kB Applications/AXUIViewService.app/AXUIViewService\n-rw-rw-r-- 2020-02-15T02:20:25-05:00 621 B  Applications/AXUIViewService.app/AXUIViewService-Entitlements.plist\n-rw-rw-r-- 2020-02-15T02:20:26-05:00 22 kB  Applications/AXUIViewService.app/Assets.car\n-rw-rw-r-- 2020-02-15T02:20:26-05:00 1.5 kB Applications/AXUIViewService.app/Info.plist\n-rw-rw-r-- 2020-02-15T02:20:26-05:00 8 B    Applications/AXUIViewService.app/PkgInfo\n")),(0,a.kt)("p",null,"See if ",(0,a.kt)("inlineCode",{parentName:"p"},"dyld")," is in the OTA files"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},'\u276f ipsw ota OTA.zip | grep dyld\n   \u2022 Listing files...\n-rwxr-xr-x 2020-02-15T02:22:01-05:00 1.7 GB System/Library/Caches/com.apple."dyld/dyld"_shared_cache_arm64e\n-rwxr-xr-x 2020-02-15T02:24:08-05:00 721 kB usr/lib/"dyld"\n')),(0,a.kt)("h4",{id:"extract-files-from-ota-payloads"},"Extract file(s) from OTA payloads"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"\u276f ipsw ota test-caches/30f4510f7fa8e1ecfb8d137f6081a8691cfc28b5.zip '^System/Library/.*/dyld_shared_cache.*$'\n   \u2022 Extracting ^System/Library/.*/dyld_shared_cache.*$...\n      \u2022 Extracting -rwxr-xr-x   1.5 GB  /System/Library/Caches/com.apple.dyld/dyld_shared_cache_arm64e to iPhone14,2_D63AP_19C5026i/dyld_shared_cache_arm64e\n      \u2022 Extracting -rwxr-xr-x   787 MB  /System/Library/Caches/com.apple.dyld/dyld_shared_cache_arm64e.1 to iPhone14,2_D63AP_19C5026i/dyld_shared_cache_arm64e.1\n      \u2022 Extracting -rwxr-xr-x   480 MB  /System/Library/Caches/com.apple.dyld/dyld_shared_cache_arm64e.symbols to iPhone14,2_D63AP_19C5026i/dyld_shared_cache_arm64e.symbols\n")),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"NOTE:")," you can supply a regex to match ",(0,a.kt)("em",{parentName:"p"},"(see ",(0,a.kt)("inlineCode",{parentName:"em"},"re_format(7)"),")")))}d.isMDXComponent=!0}}]);