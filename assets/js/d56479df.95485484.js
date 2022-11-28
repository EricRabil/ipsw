"use strict";(self.webpackChunkdocumentation=self.webpackChunkdocumentation||[]).push([[5059],{3905:(e,t,r)=>{r.d(t,{Zo:()=>c,kt:()=>m});var n=r(7294);function i(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function p(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function a(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?p(Object(r),!0).forEach((function(t){i(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):p(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,i=function(e,t){if(null==e)return{};var r,n,i={},p=Object.keys(e);for(n=0;n<p.length;n++)r=p[n],t.indexOf(r)>=0||(i[r]=e[r]);return i}(e,t);if(Object.getOwnPropertySymbols){var p=Object.getOwnPropertySymbols(e);for(n=0;n<p.length;n++)r=p[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(i[r]=e[r])}return i}var o=n.createContext({}),s=function(e){var t=n.useContext(o),r=t;return e&&(r="function"==typeof e?e(t):a(a({},t),e)),r},c=function(e){var t=s(e.components);return n.createElement(o.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var r=e.components,i=e.mdxType,p=e.originalType,o=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),u=s(r),m=i,w=u["".concat(o,".").concat(m)]||u[m]||d[m]||p;return r?n.createElement(w,a(a({ref:t},c),{},{components:r})):n.createElement(w,a({ref:t},c))}));function m(e,t){var r=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var p=r.length,a=new Array(p);a[0]=u;var l={};for(var o in t)hasOwnProperty.call(t,o)&&(l[o]=t[o]);l.originalType=e,l.mdxType="string"==typeof e?e:i,a[1]=l;for(var s=2;s<p;s++)a[s]=r[s];return n.createElement.apply(null,a)}return n.createElement.apply(null,r)}u.displayName="MDXCreateElement"},2730:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>o,contentTitle:()=>a,default:()=>d,frontMatter:()=>p,metadata:()=>l,toc:()=>s});var n=r(7462),i=(r(7294),r(3905));const p={id:"wallpaper",title:"wallpaper",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"wallpaper",description:"Dump wallpaper as PNG",last_update:{date:new Date("2022-11-28T19:49:26.000Z"),author:"blacktop"}},a=void 0,l={unversionedId:"cli/ipsw/idev/springb/wallpaper",id:"cli/ipsw/idev/springb/wallpaper",title:"wallpaper",description:"Dump wallpaper as PNG",source:"@site/docs/cli/ipsw/idev/springb/wallpaper.md",sourceDirName:"cli/ipsw/idev/springb",slug:"/cli/ipsw/idev/springb/wallpaper",permalink:"/ipsw/docs/cli/ipsw/idev/springb/wallpaper",draft:!1,editUrl:"https://github.com/blacktop/ipsw/tree/master/www/docs/cli/ipsw/idev/springb/wallpaper.md",tags:[],version:"current",frontMatter:{id:"wallpaper",title:"wallpaper",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"wallpaper",description:"Dump wallpaper as PNG",last_update:{date:"2022-11-28T19:49:26.000Z",author:"blacktop"}},sidebar:"cli",previous:{title:"orient",permalink:"/ipsw/docs/cli/ipsw/idev/springb/orient"},next:{title:"syslog",permalink:"/ipsw/docs/cli/ipsw/idev/syslog"}},o={},s=[{value:"ipsw idev springb wallpaper",id:"ipsw-idev-springb-wallpaper",level:2},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],c={toc:s};function d(e){let{components:t,...r}=e;return(0,i.kt)("wrapper",(0,n.Z)({},c,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"ipsw-idev-springb-wallpaper"},"ipsw idev springb wallpaper"),(0,i.kt)("p",null,"Dump wallpaper as PNG"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"ipsw idev springb wallpaper [flags]\n")),(0,i.kt)("h3",{id:"options"},"Options"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"  -h, --help            help for wallpaper\n  -o, --output string   Folder to save wallpaper\n")),(0,i.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre"},"      --color           colorize output\n      --config string   config file (default is $HOME/.ipsw.yaml)\n  -u, --udid string     Device UniqueDeviceID to connect to\n  -V, --verbose         verbose output\n")),(0,i.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/idev/springb"},"ipsw idev springb"),"\t - SpringBoard commands")))}d.isMDXComponent=!0}}]);