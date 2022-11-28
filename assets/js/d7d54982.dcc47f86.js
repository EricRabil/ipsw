"use strict";(self.webpackChunkdocumentation=self.webpackChunkdocumentation||[]).push([[7084],{3905:(e,t,a)=>{a.d(t,{Zo:()=>o,kt:()=>m});var l=a(7294);function d(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function r(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);t&&(l=l.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,l)}return a}function i(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?r(Object(a),!0).forEach((function(t){d(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):r(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function n(e,t){if(null==e)return{};var a,l,d=function(e,t){if(null==e)return{};var a,l,d={},r=Object.keys(e);for(l=0;l<r.length;l++)a=r[l],t.indexOf(a)>=0||(d[a]=e[a]);return d}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(l=0;l<r.length;l++)a=r[l],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(d[a]=e[a])}return d}var s=l.createContext({}),p=function(e){var t=l.useContext(s),a=t;return e&&(a="function"==typeof e?e(t):i(i({},t),e)),a},o=function(e){var t=p(e.components);return l.createElement(s.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return l.createElement(l.Fragment,{},t)}},y=l.forwardRef((function(e,t){var a=e.components,d=e.mdxType,r=e.originalType,s=e.parentName,o=n(e,["components","mdxType","originalType","parentName"]),y=p(a),m=d,u=y["".concat(s,".").concat(m)]||y[m]||c[m]||r;return a?l.createElement(u,i(i({ref:t},o),{},{components:a})):l.createElement(u,i({ref:t},o))}));function m(e,t){var a=arguments,d=t&&t.mdxType;if("string"==typeof e||d){var r=a.length,i=new Array(r);i[0]=y;var n={};for(var s in t)hasOwnProperty.call(t,s)&&(n[s]=t[s]);n.originalType=e,n.mdxType="string"==typeof e?e:d,i[1]=n;for(var p=2;p<r;p++)i[p]=a[p];return l.createElement.apply(null,i)}return l.createElement.apply(null,a)}y.displayName="MDXCreateElement"},5110:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>s,contentTitle:()=>i,default:()=>c,frontMatter:()=>r,metadata:()=>n,toc:()=>p});var l=a(7462),d=(a(7294),a(3905));const r={id:"dyld",title:"dyld",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"dyld",description:"Parse dyld_shared_cache",last_update:{date:new Date("2022-11-28T19:49:26.000Z"),author:"blacktop"}},i=void 0,n={unversionedId:"cli/ipsw/dyld/dyld",id:"cli/ipsw/dyld/dyld",title:"dyld",description:"Parse dyld_shared_cache",source:"@site/docs/cli/ipsw/dyld/dyld.md",sourceDirName:"cli/ipsw/dyld",slug:"/cli/ipsw/dyld/",permalink:"/ipsw/docs/cli/ipsw/dyld/",draft:!1,editUrl:"https://github.com/blacktop/ipsw/tree/master/www/docs/cli/ipsw/dyld/dyld.md",tags:[],version:"current",frontMatter:{id:"dyld",title:"dyld",hide_title:!0,hide_table_of_contents:!0,sidebar_label:"dyld",description:"Parse dyld_shared_cache",last_update:{date:"2022-11-28T19:49:26.000Z",author:"blacktop"}},sidebar:"cli",previous:{title:"dtree",permalink:"/ipsw/docs/cli/ipsw/dtree"},next:{title:"a2f",permalink:"/ipsw/docs/cli/ipsw/dyld/a2f"}},s={},p=[{value:"ipsw dyld",id:"ipsw-dyld",level:2},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],o={toc:p};function c(e){let{components:t,...a}=e;return(0,d.kt)("wrapper",(0,l.Z)({},o,a,{components:t,mdxType:"MDXLayout"}),(0,d.kt)("h2",{id:"ipsw-dyld"},"ipsw dyld"),(0,d.kt)("p",null,"Parse dyld_shared_cache"),(0,d.kt)("pre",null,(0,d.kt)("code",{parentName:"pre"},"ipsw dyld [flags]\n")),(0,d.kt)("h3",{id:"options"},"Options"),(0,d.kt)("pre",null,(0,d.kt)("code",{parentName:"pre"},"  -h, --help   help for dyld\n")),(0,d.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,d.kt)("pre",null,(0,d.kt)("code",{parentName:"pre"},"      --color           colorize output\n      --config string   config file (default is $HOME/.ipsw.yaml)\n  -V, --verbose         verbose output\n")),(0,d.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,d.kt)("ul",null,(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw"},"ipsw"),"\t - Download and Parse IPSWs (and SO much more)"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/a2f"},"ipsw dyld a2f"),"\t - Lookup function containing unslid address"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/a2o"},"ipsw dyld a2o"),"\t - Convert dyld_shared_cache address to offset"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/a2s"},"ipsw dyld a2s"),"\t - Lookup symbol at unslid address"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/disass"},"ipsw dyld disass"),"\t - Disassemble dyld_shared_cache at symbol/vaddr"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/dump"},"ipsw dyld dump"),"\t - Dump dyld_shared_cache data at given virtual address"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/extract"},"ipsw dyld extract"),"\t - Extract dyld_shared_cache from DMG in IPSW"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/image"},"ipsw dyld image"),"\t - Dump image array info"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/imports"},"ipsw dyld imports"),"\t - List all dylibs that load a given dylib"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/info"},"ipsw dyld info"),"\t - Parse dyld_shared_cache"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/macho"},"ipsw dyld macho"),"\t - Parse a dylib file"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/o2a"},"ipsw dyld o2a"),"\t - Convert dyld_shared_cache offset to address"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/objc"},"ipsw dyld objc"),"\t - Dump Objective-C Optimization Info"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/patches"},"ipsw dyld patches"),"\t - Dump dyld patch info"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/slide"},"ipsw dyld slide"),"\t - Dump slide info"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/split"},"ipsw dyld split"),"\t - Extracts all the dyld_shared_cache libraries"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/str"},"ipsw dyld str"),"\t - Search dyld_shared_cache for string"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/symaddr"},"ipsw dyld symaddr"),"\t - Lookup or dump symbol(s)"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/tbd"},"ipsw dyld tbd"),"\t - Generate a .tbd file for a dylib"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/webkit"},"ipsw dyld webkit"),"\t - Get WebKit version from a dyld_shared_cache"),(0,d.kt)("li",{parentName:"ul"},(0,d.kt)("a",{parentName:"li",href:"/docs/cli/ipsw/dyld/xref"},"ipsw dyld xref"),"\t - \ud83d\udea7 ","[WIP]"," Find all cross references to an address")))}c.isMDXComponent=!0}}]);