"use strict";(globalThis["webpackChunkfix_panel"]=globalThis["webpackChunkfix_panel"]||[]).push([[171],{9171:(e,t,a)=>{a.r(t),a.d(t,{default:()=>Ye});var s=a(9835),l=a(6970),n=a(1957),i=a(970),o=a.n(i),r=a(7161),c=a.n(r);const u=(0,s._)("div",{class:"col text-center text-weight-bold"}," 电务器材全生命周期管理系统 ",-1),d={class:"q-pa-md q-gutter-md"},p={key:0,style:{width:"32%","z-index":"2",position:"absolute",left:"60%"}},m={class:"col-7 q-gutter-x-md"},_={class:"row q-gutter-xs"},w={class:"col-md-12 text-h6",style:{"font-weight":"bold"}},f={class:"row"},h={class:"col-md-12"},g={class:"col-col-md-1"},y={class:"row"},b={class:"col-12 text-h6"},v={key:1,style:{width:"140px",height:"220px","z-index":"2",position:"absolute",left:"92.6%"}},k=(0,s._)("img",{src:o()},null,-1),x={class:"text-h5",style:{"text-align":"center","margin-top":"-20%"}},W={class:"text-h5",style:{"text-align":"center","margin-top":"-20%"}},z={class:"row q-col-gutter-x-md",style:{"margin-top":"15px"}},q={class:"col-md-12"},E={class:"text-h6"},D={class:"row q-col-gutter-x-md"},C={class:"col-md-7"},U={class:"row"},I={class:"col-md-6"},$={class:"col-md-6"},M={class:"col-md-5"},O={class:"row"},T={class:"col-md-7"},N={class:"col-md-5"},Z={class:"row q-col-gutter-x-md",style:{"margin-top":"15px"}},P={class:"col-3"},F={class:"row"},S={class:"col-md-12"},H=(0,s._)("img",{src:c()},null,-1),Q={class:"row",style:{"margin-top":"20px"}},B={class:"col-md-12"},V={class:"text-h6"},A={class:"col-6"},j={class:"row q-col-gutter-x-md"},Y={class:"col-6"},X=(0,s._)("div",{class:"text-h6"},"日常生产任务",-1),K={class:"col-6"},J=(0,s._)("div",{class:"text-h6"},"故障修任务",-1),R={class:"row",style:{"margin-top":"20px"}},L={class:"col-md-12"},G=(0,s._)("div",{id:"myChart",style:{width:"100%",height:"315px"}},null,-1),ee={class:"col-3"},te=(0,s._)("div",{class:"text-h6"},"检修日志",-1),ae={key:0},se={key:1},le={class:"text-right text-caption text-grey"};function ne(e,t,a,i,o,r){const c=(0,s.up)("q-btn"),ne=(0,s.up)("q-badge"),ie=(0,s.up)("q-bar"),oe=(0,s.up)("q-card-section"),re=(0,s.up)("q-separator"),ce=(0,s.up)("q-card"),ue=(0,s.up)("q-avatar"),de=(0,s.up)("q-item-section"),pe=(0,s.up)("q-item"),me=(0,s.up)("q-select"),_e=(0,s.up)("q-td"),we=(0,s.up)("q-tr"),fe=(0,s.up)("q-table"),he=(0,s.up)("q-timeline-entry"),ge=(0,s.up)("q-timeline"),ye=(0,s.up)("q-page");return(0,s.wg)(),(0,s.j4)(ye,{style:{width:"100%",height:"100%"}},{default:(0,s.w5)((()=>[(0,s.Wm)(ie,{dark:"",class:"bg-primary text-white",style:{height:"50px"}},{default:(0,s.w5)((()=>[u,(0,s._)("div",d,[(0,s.Wm)(c,{onClick:e.refresh,size:"md",dense:"",flat:"",icon:"refresh",style:{"margin-right":"10px"}},null,8,["onClick"]),(0,s.Wm)(c,{onClick:e.showMessageDialog,dense:"",flat:"",icon:"email",size:"15px"},{default:(0,s.w5)((()=>[e.unreadNumber?((0,s.wg)(),(0,s.j4)(ne,{key:0,color:"red",floating:""},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.unreadNumber),1)])),_:1})):(0,s.kq)("",!0)])),_:1},8,["onClick"]),(0,s.Wm)(c,{onClick:e.showloginMessage,flat:"",size:"15px"},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.userInfo?.nickname),1)])),_:1},8,["onClick"])])])),_:1}),e.showMessage?((0,s.wg)(),(0,s.iD)("div",p,[(0,s.Wm)(ce,{style:{height:"1027px",overflow:"auto"}},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,{style:{height:"40px"},class:"text-h5 q-pb-none"},{default:(0,s.w5)((()=>[(0,s.Uk)(" 检修分配任务消息 ")])),_:1}),((0,s.wg)(!0),(0,s.iD)(s.HY,null,(0,s.Ko)(e.task_message,((t,a)=>((0,s.wg)(),(0,s.j4)(oe,{onClick:a=>e.readMessage(t.uuid),key:a},{default:(0,s.w5)((()=>[(0,s._)("div",{class:(0,l.C_)(t.sign?"row":"row text-grey-5")},[(0,s._)("div",m,[(0,s._)("div",_,[(0,s._)("div",w,(0,l.zw)(t.BusinessText),1),(0,s._)("div",f,[(0,s._)("div",h,(0,l.zw)(t.content),1)])])]),(0,s._)("div",g,[(0,s._)("div",y,[(0,s._)("div",b,(0,l.zw)(t.created_time),1)])])],2),(0,s.Wm)(re)])),_:2},1032,["onClick"])))),128))])),_:1})])):(0,s.kq)("",!0),e.loginOut?((0,s.wg)(),(0,s.iD)("div",v,[(0,s.Wm)(ce,null,{default:(0,s.w5)((()=>[(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(ue,{color:"primary",size:"100px"},{default:(0,s.w5)((()=>[k])),_:1})])),_:1}),(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[(0,s._)("div",x,(0,l.zw)(e.userInfo?.nickname),1)])),_:1}),(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[(0,s._)("div",W,[(0,s.Wm)(c,{onClick:e.goLoginOut,color:"primary",rounded:"",label:"退出登录"},null,8,["onClick"])])])),_:1})])),_:1})])):(0,s.kq)("",!0),(0,s._)("div",z,[(0,s._)("div",q,[(0,s.Wm)(ce,{style:{width:"100%",height:"100%"}},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,{style:{height:"50px"},class:"text-h6 q-pb-none"},{default:(0,s.w5)((()=>[(0,s._)("div",E,[(0,s.Uk)(" 搜索 "),(0,s.Wm)(c,{onClick:e.search,label:"搜索",color:"primary",style:{float:"right","margin-top":"-9px"},icon:"search"},null,8,["onClick"])])])),_:1}),(0,s.Wm)(oe,{style:{height:"50px"}},{default:(0,s.w5)((()=>[(0,s._)("div",D,[(0,s._)("div",C,[(0,s._)("div",U,[(0,s._)("div",I,[(0,s.Wm)(me,{style:{"margin-top":"-15px"},dense:"","options-dense":"",outlined:"",modelValue:e.category_unique_code,"onUpdate:modelValue":t[0]||(t[0]=t=>e.category_unique_code=t),options:e.kindOptions,"use-input":"","hide-selected":"",clearable:"","fill-input":"","input-debounce":"0","map-options":"","emit-value":"","option-label":"name","option-value":"unique_code",onFilter:e.filterCat,label:"种类"},{"no-option":(0,s.w5)((()=>[(0,s.Wm)(pe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(de,{class:"text-grey"},{default:(0,s.w5)((()=>[(0,s.Uk)(" 未找到 ")])),_:1})])),_:1})])),_:1},8,["modelValue","options","onFilter"])]),(0,s._)("div",$,[(0,s.Wm)(me,{style:{"margin-top":"-15px"},dense:"","options-dense":"",outlined:"",modelValue:e.model_unique_code,"onUpdate:modelValue":t[1]||(t[1]=t=>e.model_unique_code=t),options:e.typeOptions,"use-input":"","hide-selected":"",clearable:"","fill-input":"","input-debounce":"0","map-options":"","emit-value":"","option-label":"name","option-value":"unique_code",onFilter:e.filterType,label:"类型"},{"no-option":(0,s.w5)((()=>[(0,s.Wm)(pe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(de,{class:"text-grey"},{default:(0,s.w5)((()=>[(0,s.Uk)(" 未找到 ")])),_:1})])),_:1})])),_:1},8,["modelValue","options","onFilter"])])])]),(0,s._)("div",M,[(0,s._)("div",O,[(0,s._)("div",T,[(0,s.Wm)(me,{style:{"margin-top":"-15px"},dense:"","options-dense":"",outlined:"",modelValue:e.yearModel,"onUpdate:modelValue":t[2]||(t[2]=t=>e.yearModel=t),options:e.yearOptions,"use-input":"",clearable:"","input-debounce":"0",onFilter:e.filterYear,"map-options":"","emit-value":"",label:"年"},{"no-option":(0,s.w5)((()=>[(0,s.Wm)(pe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(de,{class:"text-grey"},{default:(0,s.w5)((()=>[(0,s.Uk)(" 未找到 ")])),_:1})])),_:1})])),_:1},8,["modelValue","options","onFilter"])]),(0,s._)("div",N,[(0,s.Wm)(me,{style:{"margin-top":"-15px"},dense:"","options-dense":"",outlined:"",modelValue:e.monthModel,"onUpdate:modelValue":t[3]||(t[3]=t=>e.monthModel=t),options:e.monthOptions,"use-input":"",clearable:"","input-debounce":"0",onFilter:e.filterMonth,"map-options":"","emit-value":"",label:"月"},{"no-option":(0,s.w5)((()=>[(0,s.Wm)(pe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(de,{class:"text-grey"},{default:(0,s.w5)((()=>[(0,s.Uk)(" 未找到 ")])),_:1})])),_:1})])),_:1},8,["modelValue","options","onFilter"])])])])])])),_:1})])),_:1})])]),(0,s._)("div",Z,[(0,s._)("div",P,[(0,s._)("div",F,[(0,s._)("div",S,[(0,s.Wm)(ce,{style:{width:"100%",height:"100%"},flat:""},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,{style:{height:"40px"},class:"text-h6 q-pb-none"},{default:(0,s.w5)((()=>[(0,s.Uk)(" 作业人员主页 ")])),_:1}),(0,s.Wm)(oe,{style:{float:"left"}},{default:(0,s.w5)((()=>[(0,s.Wm)(ue,{size:"100px"},{default:(0,s.w5)((()=>[H])),_:1})])),_:1}),(0,s.Wm)(oe,{style:{height:"150px"}},{default:(0,s.w5)((()=>[(0,s.Wm)(pe,{dense:"",class:"q-px-md ellipsis"},{default:(0,s.w5)((()=>[(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)("姓名:")])),_:1}),(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.userInfo?.nickname),1)])),_:1})])),_:1}),(0,s.Wm)(pe,{dense:"",class:"q-px-md ellipsis"},{default:(0,s.w5)((()=>[(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)("联系电话:")])),_:1}),(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.userInfo?.phone),1)])),_:1})])),_:1}),(0,s.Wm)(pe,{dense:"",class:"q-px-md ellipsis"},{default:(0,s.w5)((()=>[(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)("所属车间:")])),_:1}),(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.userInfo?.workshop_name),1)])),_:1})])),_:1}),(0,s.Wm)(pe,{dense:"",class:"q-px-md ellipsis"},{default:(0,s.w5)((()=>[(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)("所属工区:")])),_:1}),(0,s.Wm)(de,null,{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.userInfo?.work_area_name),1)])),_:1})])),_:1})])),_:1})])),_:1})])]),(0,s._)("div",Q,[(0,s._)("div",B,[(0,s.Wm)(ce,{style:{width:"100%",height:"100%"}},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[(0,s._)("div",V,"周期修计划("+(0,l.zw)(e.planYear)+"年)",1)])),_:1}),(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(fe,{class:"my-sticky-virtscroll-table",style:{height:"594px"},"virtual-scroll":"","rows-per-page-options":[0],dense:"",bordered:"","row-key":"name",rows:e.planrows,columns:e.plancolumns},{body:(0,s.w5)((e=>[(0,s.Wm)(we,{props:e},{default:(0,s.w5)((()=>[(0,s.Wm)(_e,{key:"subModelName",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.subModelName),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"year_plan_device_count",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.year_plan_device_count),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[0]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[0]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[1]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[1]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[2]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[2]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[3]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[3]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[4]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[4]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[5]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[5]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[6]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[6]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[7]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[7]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[8]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[8]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[9]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[9]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[10]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[10]),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"statistics[11]",props:e},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(e.row.statistics[11]),1)])),_:2},1032,["props"])])),_:2},1032,["props"])])),_:1},8,["rows","columns"])])),_:1})])),_:1})])])]),(0,s._)("div",A,[(0,s._)("div",j,[(0,s._)("div",Y,[(0,s.Wm)(ce,{flat:""},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[X])),_:1}),(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(fe,{class:"my-sticky-virtscroll-table",style:{height:"400px"},"virtual-scroll":"",dense:"","rows-per-page-options":[0],rows:e.taskrows,columns:e.taskcolumns,"row-key":"name"},{body:(0,s.w5)((t=>[(0,s.Wm)(we,{props:t},{default:(0,s.w5)((()=>[(0,s.Wm)(_e,{key:"name",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.name),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"task",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.task),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"finish",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.finish),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"check",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.check),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"overhaul",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.overhaul),1)])),_:2},1032,["props"]),(0,s.Wm)(_e,{key:"operation",props:t},{default:(0,s.w5)((()=>[(0,s.Wm)(c,{onClick:a=>e.finish(t.row.operation),flat:"",size:"17px",color:"primary",label:"提交"},null,8,["onClick"])])),_:2},1032,["props"])])),_:2},1032,["props"])])),_:1},8,["rows","columns"])])),_:1})])),_:1})]),(0,s._)("div",K,[(0,s.Wm)(ce,{flat:""},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[J])),_:1}),(0,s.Wm)(oe,null,{default:(0,s.w5)((()=>[(0,s.Wm)(fe,{class:"my-sticky-virtscroll-table",style:{height:"400px"},"virtual-scroll":"","wrap-cells":"","rows-per-page-options":[0],rows:e.faultrows,columns:e.faultcolumns,"row-key":"name","binary-state-sort":""},{body:(0,s.w5)((t=>[(0,s.Wm)(we,{props:t},{default:(0,s.w5)((()=>[(0,s.Wm)(_e,{key:"code",class:(0,l.C_)("EXTENDED"==t.row.operation.status?"text-red-6":""),props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(`${t.row.code} ${t.row.type}`),1)])),_:2},1032,["class","props"]),(0,s.Wm)(_e,{key:"expirationDate",class:(0,l.C_)("EXTENDED"==t.row.operation.status?"text-red-6":""),props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.expirationDate),1)])),_:2},1032,["class","props"]),(0,s.wy)((0,s.Wm)(_e,{key:"completeDate",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.completeDate),1)])),_:2},1032,["props"]),[[n.F8,"COMMITTED"==t.row.operation.status]]),(0,s.wy)((0,s.Wm)(_e,{key:"completeDate",props:t},{default:(0,s.w5)((()=>["UNDONE"==t.row.operation.status||"EXTENDED"==t.row.operation.status?((0,s.wg)(),(0,s.j4)(c,{key:0,flat:"",class:"full-width",onClick:a=>e.commitFix(t.row.operation.uuid),size:"17px",color:"EXTENDED"==t.row.operation.status?"red":"primary",label:"提交"},null,8,["onClick","color"])):(0,s.kq)("",!0)])),_:2},1032,["props"]),[[n.F8,"UNDONE"==t.row.operation.status||"EXTENDED"==t.row.operation.status]])])),_:2},1032,["props"])])),_:1},8,["rows","columns"])])),_:1})])),_:1})])]),(0,s._)("div",R,[(0,s._)("div",L,[(0,s.Wm)(ce,{style:{width:"100%",height:"100%"},flat:""},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,{style:{height:"52px"},class:"text-h6 q-pb-none"},{default:(0,s.w5)((()=>[(0,s.Uk)(" 任务统计（日常生产任务） ")])),_:1}),(0,s.Wm)(oe,{class:"text-h6 q-pb-none"},{default:(0,s.w5)((()=>[G])),_:1})])),_:1})])])]),(0,s._)("div",ee,[(0,s.Wm)(ce,{style:{height:"100%",width:"100%"},flat:""},{default:(0,s.w5)((()=>[(0,s.Wm)(oe,{class:"row q-pb-none"},{default:(0,s.w5)((()=>[te])),_:1}),(0,s.Wm)(oe,{style:{height:"800px","overflow-y":"auto"}},{default:(0,s.w5)((()=>[(0,s.Wm)(ge,{color:"primary",style:{"overflow-y":"auto"}},{default:(0,s.w5)((()=>[((0,s.wg)(!0),(0,s.iD)(s.HY,null,(0,s.Ko)(e.log,((t,a)=>((0,s.wg)(),(0,s.j4)(he,{key:a},{title:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.name),1)])),default:(0,s.w5)((()=>[(0,s._)("div",null,(0,l.zw)(t.unique_code),1),(0,s._)("div",null,(0,l.zw)(t.info),1),(0,s._)("div",null,(0,l.zw)(t.description),1),null!==t?.fix_workflow_process&&"JSON2"==t?.fix_workflow_process?.check_type&&t?.fix_workflow_process?.values?((0,s.wg)(),(0,s.iD)("div",ae,[(0,s.Wm)(fe,{class:"my-sticky-virtscroll-table",style:{height:"300px"},"virtual-scroll":"","hide-bottom":"",dense:"","rows-per-page-options":[0],rows:t.testrows,columns:e.testcolumns,"row-key":"name"},null,8,["rows","columns"])])):(0,s.kq)("",!0),null!==t?.fix_workflow_process&&"PDF"==t?.fix_workflow_process?.check_type?((0,s.wg)(),(0,s.iD)("div",se,[(0,s.Wm)(c,{onClick:a=>e.lookPDF(t.url),color:"primary",label:"点击查看PDF"},null,8,["onClick"])])):(0,s.kq)("",!0),(0,s._)("div",le,(0,l.zw)(t.normaltime),1)])),_:2},1024)))),128))])),_:1})])),_:1})])),_:1})])])])),_:1})}a(9665);var ie=a(7406),oe=a(3632),re=a.n(oe),ce=a(499),ue=a(9302),de=a(6827);const pe={class:"text-h6"},me={class:"text-h6"};function _e(e,t,a,n,i,o){const r=(0,s.up)("q-card-section"),c=(0,s.up)("q-checkbox"),u=(0,s.up)("q-td"),d=(0,s.up)("q-btn"),p=(0,s.up)("q-tr"),m=(0,s.up)("q-table"),_=(0,s.up)("q-card-actions"),w=(0,s.up)("q-card"),f=(0,s.up)("q-dialog");return(0,s.wg)(),(0,s.j4)(f,{ref:"dialogRef",onHide:e.onDialogHide,"full-width":"",persistent:""},{default:(0,s.w5)((()=>[(0,s.Wm)(w,{class:"q-dialog-plugin"},{default:(0,s.w5)((()=>[(0,s.Wm)(r,null,{default:(0,s.w5)((()=>[(0,s.Wm)(r,null,{default:(0,s.w5)((()=>[(0,s._)("div",pe,"已完成验收 (共 "+(0,l.zw)(e.finishrows.length)+" 条)",1)])),_:1}),(0,s.Wm)(m,{class:"my-sticky-virtscroll-table",style:{height:"400px"},"virtual-scroll":"","rows-per-page-options":[0],rows:e.finishrows,columns:e.finishcolumns,"row-key":"name"},{body:(0,s.w5)((a=>[(0,s.Wm)(p,{props:a},{default:(0,s.w5)((()=>[(0,s.Wm)(u,{key:"code",props:a},{default:(0,s.w5)((()=>[(0,s.Wm)(c,{modelValue:e.selection,"onUpdate:modelValue":t[0]||(t[0]=t=>e.selection=t),val:a.row.code,color:"grey-8"},null,8,["modelValue","val"]),(0,s.Uk)(" "+(0,l.zw)(a.row.code),1)])),_:2},1032,["props"]),(0,s.Wm)(u,{key:"type",props:a},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(a.row.type),1)])),_:2},1032,["props"]),(0,s.Wm)(u,{key:"overhaultime",props:a},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(a.row.overhaultime),1)])),_:2},1032,["props"]),(0,s.Wm)(u,{key:"checktime",props:a},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(a.row.checktime),1)])),_:2},1032,["props"]),(0,s.Wm)(u,{key:"operation",props:a},{default:(0,s.w5)((()=>[(0,s.Wm)(d,{onClick:t=>e.deleteFinishData(a.row.operation),flat:"",size:"18px",color:"primary",label:"删除"},null,8,["onClick"])])),_:2},1032,["props"])])),_:2},1032,["props"])])),_:1},8,["rows","columns"]),(0,s.Wm)(_,{align:"right"},{default:(0,s.w5)((()=>[(0,s.Wm)(d,{color:"primary",outline:"",label:"关闭",size:"18px",onClick:e.onCancelClick},null,8,["onClick"]),(0,s.Wm)(d,{color:"primary",label:"提交",size:"18px",onClick:e.onOKClick},null,8,["onClick"])])),_:1})])),_:1}),(0,s.Wm)(r,null,{default:(0,s.w5)((()=>[(0,s.Wm)(r,null,{default:(0,s.w5)((()=>[(0,s._)("div",me,"未完成验收 (共 "+(0,l.zw)(e.unfinishrows.length)+" 条)",1)])),_:1}),(0,s.Wm)(m,{class:"my-sticky-virtscroll-table",style:{height:"400px"},"virtual-scroll":"","rows-per-page-options":[0],rows:e.unfinishrows,columns:e.unfinishcolumns,"row-key":"name"},{body:(0,s.w5)((t=>[(0,s.Wm)(p,{props:t},{default:(0,s.w5)((()=>[(0,s.Wm)(u,{key:"code",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.code),1)])),_:2},1032,["props"]),(0,s.Wm)(u,{key:"type",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.type),1)])),_:2},1032,["props"]),(0,s.Wm)(u,{key:"overhaultime",props:t},{default:(0,s.w5)((()=>[(0,s.Uk)((0,l.zw)(t.row.overhaultime),1)])),_:2},1032,["props"]),(0,s.Wm)(u,{key:"operation",props:t},{default:(0,s.w5)((()=>[(0,s.Wm)(d,{onClick:a=>e.deleteUnfinishData(t.row.operation),size:"19px",flat:"",color:"primary",label:"删除"},null,8,["onClick"])])),_:2},1032,["props"])])),_:2},1032,["props"])])),_:1},8,["rows","columns"])])),_:1})])),_:1})])),_:1},8,["onHide"])}var we=a(1569),fe=a(906);const he=(0,s.aZ)({props:{fixer_id:Number,operation:Object},emits:[...fe.Z.emits],setup(e){const t=(0,ue.Z)(),a=e.fixer_id,l=e.operation.uuid,n=e.operation.sub_model_unique_code;let i=[],o=[];const r=[{name:"code",required:!0,label:"唯一编号",align:"center",field:"code",style:"font-size:18px"},{name:"type",align:"center",label:"种类型",field:"type",style:"font-size:18px;white-space:pre;"},{name:"overhaultime",align:"center",label:"检修时间",field:"overhaultime",style:"font-size:18px",sortable:!0,sort:(e,t)=>new Date(t).getTime()-new Date(e).getTime()},{name:"checktime",align:"center",label:"验收时间",field:"checktime",style:"font-size:18px",sortable:!0,sort:(e,t)=>new Date(t).getTime()-new Date(e).getTime()},{name:"operation",align:"center",label:"操作",field:"operation"}],c=(0,ce.iH)([]),u=(0,ce.iH)([]),d=[{name:"code",required:!0,label:"唯一编号",align:"center",field:"code",style:"font-size:18px"},{name:"type",align:"center",label:"种类型",field:"type",style:"font-size:18px;white-space:pre;"},{name:"overhaultime",align:"center",label:"检修时间",field:"overhaultime",style:"font-size:18px",sortable:!0,sort:(e,t)=>new Date(t).getTime()-new Date(e).getTime()},{name:"operation",align:"center",label:"操作",field:"operation"}],p=(0,ce.iH)([]),{dialogRef:m,onDialogHide:_,onDialogOK:w,onDialogCancel:f}=(0,fe.Z)(),h=()=>{we.apis.finishData(["EntireInstance","EntireInstance.SubModel","EntireInstance.SubModel.Parent","EntireInstance.SubModel.Parent.Category"],1,a,"checked_at desc",0,0,n).then((e=>{c.value=[],i=[],e.data.content.repair_base_normal_fix_complete_entire_instances.forEach((e=>{const{fixed_at:t,checked_at:a}=e,s=t.substring(0,10),l=s,n=a.substring(0,10),o=n;i.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,overhaultime:`${l}`,checktime:`${o}`,operation:`${e.uuid}`}),c.value=i}))}))},g=()=>{we.apis.finishData(["EntireInstance","EntireInstance.SubModel","EntireInstance.SubModel.Parent","EntireInstance.SubModel.Parent.Category"],0,a,"fixed_at desc",0,0,n).then((e=>{p.value=[],o=[],e.data.content.repair_base_normal_fix_complete_entire_instances.forEach((e=>{const{fixed_at:t}=e,a=t.substring(0,10),s=a;o.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,overhaultime:`${s}`,operation:`${e.uuid}`}),p.value=o}))}))},y=e=>{t.dialog({title:"确认删除",message:"此操作会删除当前已完成验收的器材（此操作不可恢复）",ok:{label:"确认",color:"red"},cancel:{flat:!0,label:"取消"},persistent:!0}).onOk((()=>{we.apis.deleteNormalFixCompleteEntireInstance(e).then((e=>{t.notify({type:"positive",message:e.data.msg}),h()}))}))},b=e=>{t.dialog({title:"确认删除",message:"此操作会删除当前未完成验收的器材（此操作不可恢复）",ok:{label:"确认",color:"red"},cancel:{flat:!0,label:"取消"},persistent:!0}).onOk((()=>{we.apis.deleteNormalFixCompleteEntireInstance(e).then((e=>{t.notify({type:"positive",message:e.data.msg}),g()}))}))};return(0,s.wF)((()=>{h(),g()})),{finishcolumns:r,finishrows:c,selection:u,dialogRef:m,unfinishrows:p,unfinishcolumns:d,deleteFinishData:y,deleteUnfinishData:b,onDialogHide:_,onOKClick(){0!=u.value.length?w({selection:u,uuid:l}):t.dialog({title:"提示",message:"还未选择要提交的器材！",ok:{label:"关闭",color:"red"},persistent:!0})},onCancelClick:f}}});var ge=a(1639),ye=a(7743),be=a(4458),ve=a(3190),ke=a(1121),xe=a(3532),We=a(7220),ze=a(1221),qe=a(4455),Ee=a(1821),De=a(9984),Ce=a.n(De);const Ue=(0,ge.Z)(he,[["render",_e]]),Ie=Ue;Ce()(he,"components",{QDialog:ye.Z,QCard:be.Z,QCardSection:ve.Z,QTable:ke.Z,QTr:xe.Z,QTd:We.Z,QCheckbox:ze.Z,QBtn:qe.Z,QCardActions:Ee.Z});var $e=a(5317),Me=a(8339),Oe=a(6928);const Te=(0,s.aZ)({name:"Panel",setup(){const e=(0,Oe.t)(),t=(0,Me.tv)(),a=(0,ue.Z)();let l;const n=(0,ce.iH)(null);let i,o;const r=(0,ce.iH)();let c,u,d,p,m;const _=(0,ce.iH)(!1),w=(0,ce.iH)(!1),f=(0,ce.iH)([]),h=(0,ce.iH)([]);let g=[],y=[],b=[];const v=(0,ce.iH)([]);let k=[],x=[],W=[];const z=(0,ce.iH)(),q=(0,ce.iH)([]),E=(0,ce.iH)([]),D=(0,ce.iH)(0),C=(0,ce.iH)();let U=[];const I=(0,ce.iH)();let $=[];const M=(0,ce.iH)([]),O=[],T=(0,ce.iH)([]),N=[];let Z=(0,ce.iH)(null),P=(0,ce.iH)(null);const F=(0,ce.iH)(),S=(0,ce.iH)();let H=[],Q=[];const B=[{name:"subModelName",required:!0,label:"型号",align:"center",field:"subModelName",style:"font-size:15px"},{name:"year_plan_device_count",required:!0,align:"center",label:"小计",field:"year_plan_device_count",style:"font-size:15px"},{name:"statistics[0]",required:!0,align:"center",label:"1月",field:"statistics[0]",style:"font-size:15px"},{name:"statistics[1]",required:!0,align:"center",label:"2月",field:"statistics[1]",style:"font-size:15px"},{name:"statistics[2]",required:!0,align:"center",label:"3月",field:"statistics[2]",style:"font-size:15px"},{name:"statistics[3]",required:!0,align:"center",label:"4月",field:"statistics[3]",style:"font-size:15px"},{name:"statistics[4]",required:!0,align:"center",label:"5月",field:"statistics[4]",style:"font-size:15px"},{name:"statistics[5]",required:!0,align:"center",label:"6月",field:"statistics[5]",style:"font-size:15px"},{name:"statistics[6]",required:!0,align:"center",label:"7月",field:"statistics[6]",style:"font-size:15px"},{name:"statistics[7]",required:!0,align:"center",label:"8月",field:"statistics[7]",style:"font-size:15px"},{name:"statistics[8]",required:!0,align:"center",label:"9月",field:"statistics[8]",style:"font-size:15px"},{name:"statistics[9]",required:!0,align:"center",label:"10月",field:"statistics[9]",style:"font-size:15px"},{name:"statistics[10]",required:!0,align:"center",label:"11月",field:"statistics[10]",style:"font-size:15px"},{name:"statistics[11]",required:!0,align:"center",label:"12月",field:"statistics[11]",style:"font-size:15px"}],V=[{name:"name",align:"center",label:"型号",field:"type",style:"font-size:15px"},{name:"task",align:"center",label:"任务",field:"task",style:"font-size:15px"},{name:"finish",align:"center",label:"完成",field:"finish",style:"font-size:15px"},{name:"check",align:"center",label:"验收",field:"check",style:"font-size:15px"},{name:"overhaul",align:"center",label:"检修",field:"overhaul",style:"font-size:15px"},{name:"operation",align:"center",label:"操作",field:"operation"}],A=[{name:"item",align:"left",label:"测试项目",field:"item",sortable:!0},{name:"testvalue",align:"left",label:"测试值",field:"testvalue",style:"max-width: 45px",classes:"ellipsis"},{name:"standardvalue",align:"left",label:"标准值",field:"standardvalue",style:"width: 15px"},{name:"result",align:"left",label:"测试结果",field:"result",style:"width:15px"}],j=(0,ce.iH)([]),Y=(0,ce.iH)(),X=[{name:"code",required:!0,label:"唯一编号/种类型",align:"center",style:"font-size:15px;"},{name:"expirationDate",required:!0,label:"截止时间",align:"center",field:"expirationDate",style:"font-size:15px;"},{name:"completeDate",required:!0,label:"提交时间",align:"center",field:"completeDate",style:"font-size:15px;"}];z.value={legend:{},tooltip:{},dataZoom:[{show:!0,realtime:!0,start:0,end:100,xAxisIndex:[0,1]},{type:"inside",realtime:!0,start:0,end:100,xAxisIndex:[0,1]}],dataset:{source:[["product","任务","检修","验收"]]},xAxis:{type:"category"},yAxis:{},series:[{type:"bar"},{type:"bar"},{type:"bar"}]};const K=()=>{_.value=!_.value,w.value&&(w.value=!1)},J=()=>{w.value=!w.value,_.value&&(_.value=!1)},R=()=>{e.signOut(),t.replace("/login")},L=()=>{let e=new Date,t=e.getFullYear(),a=e.getMonth()+1;return{year:t,month:a}},G=(e,t,a)=>{t((()=>{C.value=U.filter((t=>t.name.indexOf(e)>-1))}))};(0,s.YP)(Z,((e,t)=>{e?we.apis.entireMode(e,0).then((e=>{$=e.data.content.entire_models,I.value=$})):we.apis.entireMode("",0).then((e=>{$=e.data.content.entire_models,I.value=$})),P.value=null}));const ee=(e,t,a)=>{t((()=>{I.value=$.filter((t=>t.name.indexOf(e)>-1))}))},te=()=>{const{year:e}=L();F.value=e;for(let t=0;t<=50;t++)O.push((e+t).toString());M.value=O},ae=(e,t)=>{t((()=>{M.value=O.filter((t=>t.indexOf(e)>-1))}))},se=()=>{const{month:e}=L();S.value=e;for(let t=1;t<=12;t++)N.push(t.toString());T.value=N},le=(e,t)=>{t((()=>{T.value=N.filter((t=>t.indexOf(e)>-1))}))};n.value=JSON.parse(window.localStorage.getItem("userInfo")),i=n.value.work_area_id,o=n.value.id;const ne=e=>{f.value=[],we.apis.cycleFixPlan(e).then((e=>{let{statistics_year:t,statistics_month1:a,statistics_month2:s,statistics_month3:l,statistics_month4:n,statistics_month5:o,statistics_month6:r,statistics_month7:c,statistics_month8:u,statistics_month9:d,statistics_month10:p,statistics_month11:m,statistics_month12:_}=e.data.content,w=[a,s,l,n,o,r,c,u,d,p,m,_],h={};""!==i&&t&&t.hasOwnProperty(i)&&t[i].hasOwnProperty("categories")&&t[i].hasOwnProperty("categories")&&re().each(t[i]["categories"],((e,t)=>{t.hasOwnProperty("subs")&&re().each(t["subs"],((e,t)=>{t.hasOwnProperty("subs")&&re().each(t["subs"],((e,t)=>{h.hasOwnProperty(e)||(h[e]={subModelName:t["name"],year_plan_device_count:0,statistics:[0,0,0,0,0,0,0,0,0,0,0,0]}),h[e]["year_plan_device_count"]=t["statistics"]["plan_device_count"]}))}))}));for(let f=0;f<12;f++)""!==i&&w[f]&&w[f].hasOwnProperty(i)&&w[f][i].hasOwnProperty("categories")&&re().each(w[f][i]["categories"],((e,t)=>{t.hasOwnProperty("subs")&&re().each(t["subs"],((e,t)=>{t.hasOwnProperty("subs")&&re().each(t["subs"],((e,t)=>{h.hasOwnProperty(e)&&(h[e]["statistics"][f]=t["statistics"]["plan_device_count"])}))}))}));f.value=Object.values(h)}))},oe=(e,t)=>{we.apis.entireInstanceLog(2,["EntireInstance","EntireInstance.SubModel","EntireInstance.SubModel.Parent","EntireInstance.SubModel.Parent.Category","FixWorkflowProcess"],o,e,t,0,"created_at desc").then((e=>{let t=[];0===e.data.content.entire_instance_logs.length?(t.push({name:"暂无检修日志"}),q.value=t):(e.data.content.entire_instance_logs.forEach((e=>{const{name:a,description:s,created_at:l,fix_workflow_process:n}=e,i=l.substring(0,10),o=l.substring(11,19),r=i+" "+o;if(null!=n){if("JSON2"==n.check_type){let e=[];n.values&&(n.values.body.forEach((t=>{const{test_project_name:a,test_value:s,standard_value:l,conclusion:n,unit:i}=t;let o,r;o=1==n?"通过":"未通过",r=null===s?"无":null===i?s:s+`${i}`,e.push({item:`${a}`,testvalue:`${r}`,standardvalue:`${l}`,result:`${o}`})})),j.value=e)}"PDF"==n.check_type&&(Y.value=n.pdf_url)}t.push({name:a,description:s,normaltime:r,unique_code:`${e.entire_instance.identity_code}`,info:`${e.entire_instance.category_name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,fix_workflow_process:n,testrows:j.value,url:Y.value})})),q.value=t)}))},pe=e=>{a.dialog({component:$e.Z,componentProps:{pdf_url:e}}).onOk((()=>{})).onCancel((()=>{}))},me=()=>{we.apis.category("Q",0).then((e=>{U=e.data.content.categories,C.value=U}))},_e=()=>{we.apis.fixAllocationMessage(["SubModel","EntireInstance"],"","created_at desc",o,1).then((e=>{E.value=[],H=[],D.value=0,e.data.content.repair_base_fix_allocation_messages.forEach((e=>{const{icon:t,BusinessText:a,content:s,created_at:l,uuid:n,be_un_read:i}=e,o=l.substring(0,10),r=l.substring(11,19),c=o+" "+r;H.push({icon:t,BusinessText:a,content:s,created_time:c,sign:i,uuid:n}),E.value=H})),D.value+=E.value.length,we.apis.fixAllocationMessage(["SubModel","EntireInstance"],10,"created_at desc",o,0).then((e=>{Q=[],e.data.content.repair_base_fix_allocation_messages.forEach((e=>{const{icon:t,BusinessText:a,content:s,created_at:l,be_un_read:n}=e,i=l.substring(0,10),o=l.substring(11,19),r=i+" "+o;Q.push({icon:t,BusinessText:a,content:s,created_time:r,sign:n})})),Q.forEach((e=>{E.value.push(e)}))}))})),clearInterval(u),u=window.setInterval((()=>{setTimeout((()=>{we.apis.fixAllocationMessage(["SubModel","EntireInstance"],"","created_at desc",o,1).then((e=>{E.value=[],H=[],D.value=0,e.data.content.repair_base_fix_allocation_messages.forEach((e=>{const{icon:t,BusinessText:a,content:s,created_at:l,uuid:n,be_un_read:i}=e,o=l.substring(0,10),r=l.substring(11,19),c=o+" "+r;H.push({icon:t,BusinessText:a,content:s,created_time:c,sign:i,uuid:n}),E.value=H})),D.value+=E.value.length,we.apis.fixAllocationMessage(["SubModel","EntireInstance"],10,"created_at desc",o,0).then((e=>{Q=[],e.data.content.repair_base_fix_allocation_messages.forEach((e=>{const{icon:t,BusinessText:a,content:s,created_at:l,be_un_read:n}=e,i=l.substring(0,10),o=l.substring(11,19),r=i+" "+o;Q.push({icon:t,BusinessText:a,content:s,created_time:r,sign:n})})),Q.forEach((e=>{E.value.push(e)}))}))}))}),0)}),6e4)},fe=e=>{void 0!=e&&we.apis.signRead(e).then((e=>{_e()}))},he=(e,t,a,s)=>(h.value=[],g=[],z.value.dataset.source.splice(1),we.apis.normalFixAllocation(["SubModel.Parent"],e,t,a,s,o).then((e=>{e.data.content.repair_base_normal_fix_allocations&&0!==e.data.content.repair_base_normal_fix_allocations.length&&e.data.content.repair_base_normal_fix_allocations.forEach((e=>{g.push({name:e.sub_model.name,task:e.mission_number,finish:e.completed_number,check:0,overhaul:0,operation:{uuid:e.uuid,sub_model_unique_code:e.sub_model_unique_code}})}))}))),ge=()=>(clearInterval(c),c=window.setInterval((()=>{setTimeout((()=>{we.apis.normalFixCompleteEntireInstance(o).then((e=>{e.data.content.checked_statistics&&(y=e.data.content.checked_statistics,y.forEach((e=>{for(let t in g)g[t].name==e.sub_model_name&&(g[t].check=e.aggregate)}))),e.data.content.total_statistics&&(b=e.data.content.total_statistics,b.forEach((e=>{for(let t in g)g[t].name==e.sub_model_name&&(g[t].overhaul=e.aggregate)})))}))}),0)}),6e4),we.apis.normalFixCompleteEntireInstance(o).then((e=>{e.data.content.checked_statistics&&(y=e.data.content.checked_statistics,y.forEach((e=>{for(let t in g)g[t].name==e.sub_model_name&&(g[t].check=e.aggregate)}))),e.data.content.total_statistics&&(b=e.data.content.total_statistics,b.forEach((e=>{for(let t in g)g[t].name==e.sub_model_name&&(g[t].overhaul=e.aggregate)})))}))),ye=async(e,t,a,s)=>{await he(e,t,a,s),await ge(),clearTimeout(p),p=setTimeout((()=>{h.value=[],g.forEach((e=>{z.value.dataset.source.push([`${e.name}`,e.task,e.overhaul,e.check])})),l.setOption(z.value),h.value=g}),0),clearInterval(m),m=window.setInterval((()=>{setTimeout((()=>{clearTimeout(p),p=setTimeout((()=>{h.value=[],g.forEach((e=>{z.value.dataset.source.push([`${e.name}`,e.task,e.overhaul,e.check])})),l.setOption(z.value),h.value=g}),0)}),0)}),6e4)},be=()=>{we.apis.getBreakdownFixAllocation(["EntireInstance","EntireInstance.SubModel","EntireInstance.SubModel.Parent","EntireInstance.SubModel.Parent.Category","Fixer"],o).then((e=>{v.value=[],k=[],x=[],W=[],e.data.content.repair_base_breakdown_fix_allocations.forEach((e=>{const{deadline_at:t,committed_at:a,created_at:s}=e,l=t.substring(0,10),n=l,i=s.substring(0,10),o=s.substring(11,19),r=i+" "+o;if("EXTENDED"==e.status_code&&k.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,expirationDate:`${n}`,operation:{status:e.status_code,uuid:e.uuid},sort:r}),"UNDONE"==e.status_code&&x.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,expirationDate:`${n}`,operation:{status:e.status_code,uuid:e.uuid},sort:r}),"COMMITTED"==e.status_code){const t=a.substring(0,10),s=t;W.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,expirationDate:`${n}`,completeDate:`${s}`,operation:{status:e.status_code,uuid:e.uuid},sort:r})}})),k=k.sort(((e,t)=>e.sort.localeCompare(t.sort))),x=x.sort(((e,t)=>e.sort.localeCompare(t.sort))),W=W.sort(((e,t)=>e.sort.localeCompare(t.sort))),v.value=[...k,...x,...W]})),clearInterval(d),d=window.setInterval((()=>{setTimeout((()=>{we.apis.getBreakdownFixAllocation(["EntireInstance","EntireInstance.SubModel","EntireInstance.SubModel.Parent","EntireInstance.SubModel.Parent.Category","Fixer"],o).then((e=>{v.value=[],k=[],x=[],W=[],e.data.content.repair_base_breakdown_fix_allocations.forEach((e=>{const{deadline_at:t,committed_at:a}=e,s=t.substring(0,10),l=s;if("EXTENDED"==e.status_code&&k.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,expirationDate:l,operation:{status:e.status_code,uuid:e.uuid}}),"UNDONE"==e.status_code&&x.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,expirationDate:l,operation:{status:e.status_code,uuid:e.uuid}}),"COMMITTED"==e.status_code){const t=a.substring(0,10),s=t;W.push({code:`${e.entire_instance_identity_code}`,type:`${e.entire_instance.sub_model.parent.category.name}  ${e.entire_instance.sub_model.parent.name}  ${e.entire_instance.sub_model.name}`,expirationDate:l,completeDate:`${s}`,operation:{status:e.status_code,uuid:e.uuid}})}})),k=k.sort(((e,t)=>e.expirationDate.localeCompare(t.expirationDate))),x=x.sort(((e,t)=>e.expirationDate.localeCompare(t.expirationDate))),W=W.sort(((e,t)=>e.expirationDate.localeCompare(t.expirationDate))),v.value=[...k,...x,...W]}))}),0)}),6e4)},ve=e=>{we.apis.getCommit(e).then((e=>{a.notify({type:"positive",message:e.data.msg}),be()}))},ke=()=>{r.value=Number(F.value),ne(Number(F.value)),oe(Number(F.value),Number(S.value)),ye(Z.value,P.value,Number(F.value),Number(S.value))},xe=()=>{t.go(0)},We=e=>{a.dialog({component:Ie,componentProps:{fixer_id:o,operation:e}}).onOk((({selection:e,uuid:t})=>{we.apis.commit(e.value,t).then((e=>{de.Z.create({type:"positive",message:e.data.msg.substring(0,8)}),ye(Z.value,P.value,Number(F.value),Number(S.value))}))})).onCancel((()=>{}))};return(0,s.bv)((()=>{l=ie.S1(document.getElementById("myChart")),te(),se(),r.value=F.value,ne(F.value),oe(F.value,S.value),me(),_e(),ye("","",F.value,S.value),be()})),(0,s.Jd)((()=>{clearInterval(u),clearInterval(c),clearInterval(d)})),{userInfo:n,loginOut:_,goLoginOut:R,showMessage:w,showloginMessage:K,showMessageDialog:J,log:q,planYear:r,task_message:E,unreadNumber:D,plancolumns:B,planrows:f,taskcolumns:V,taskrows:h,faultcolumns:X,faultrows:v,testcolumns:A,testrows:j,kindOptions:C,category_unique_code:Z,model_unique_code:P,yearModel:F,monthModel:S,typeOptions:I,yearOptions:M,monthOptions:T,filterCat:G,filterType:ee,filterYear:ae,filterMonth:le,search:ke,readMessage:fe,finish:We,lookPDF:pe,refresh:xe,commitFix:ve}}});var Ne=a(9885),Ze=a(6408),Pe=a(990),Fe=a(1357),Se=a(926),He=a(3587),Qe=a(490),Be=a(1233),Ve=a(6292),Ae=a(8767);const je=(0,ge.Z)(Te,[["render",ne]]),Ye=je;Ce()(Te,"components",{QPage:Ne.Z,QBar:Ze.Z,QBtn:qe.Z,QBadge:Pe.Z,QCard:be.Z,QCardSection:ve.Z,QAvatar:Fe.Z,QSeparator:Se.Z,QSelect:He.Z,QItem:Qe.Z,QItemSection:Be.Z,QTable:ke.Z,QTr:xe.Z,QTd:We.Z,QTimeline:Ve.Z,QTimelineEntry:Ae.Z})}}]);