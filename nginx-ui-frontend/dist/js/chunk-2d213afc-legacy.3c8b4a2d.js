(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d213afc"],{ae79:function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=this,a=t.$createElement,n=t._self._c||a;return n("a-card",{attrs:{title:"网站管理"}},[n("std-table",{ref:"table",attrs:{api:t.api,columns:t.columns,data_key:"configs",disable_search:!0},on:{clickEdit:function(t){return e.$router.push({path:"/domain/"+t.name})}},scopedSlots:t._u([{key:"action",fn:function(e){var a=e.record;return[a.enabled?n("a",{on:{click:function(e){return t.disable(a.name)}}},[t._v("禁用")]):n("a",{on:{click:function(e){return t.enable(a.name)}}},[t._v("启用")]),n("a-divider",{attrs:{type:"vertical"}})]}}])})],1)},s=[],o=a("d3a9"),i=[{title:"名称",dataIndex:"name",scopedSlots:{customRender:"名称"},sorter:!0,pithy:!0},{title:"状态",dataIndex:"enabled",badge:!0,scopedSlots:{customRender:"enabled"},mask:{true:"启用",false:"未启用"},sorter:!0,pithy:!0},{title:"修改时间",dataIndex:"modify",datetime:!0,scopedSlots:{customRender:"modify"},sorter:!0,pithy:!0},{title:"操作",dataIndex:"action",scopedSlots:{customRender:"action"}}],c={name:"Domain",components:{StdTable:o["a"]},data:function(){return{api:this.$api.domain,columns:i}},methods:{enable:function(e){var t=this;this.$api.domain.enable(e).then((function(){t.$message.success("启用成功"),t.$refs.table.get_list()})).catch((function(e){console.log(e),t.$message.error("启用失败")}))},disable:function(e){var t=this;this.$api.domain.disable(e).then((function(){t.$message.success("禁用成功"),t.$refs.table.get_list()})).catch((function(e){console.log(e),t.$message.error("禁用失败")}))}}},r=c,d=a("2877"),l=Object(d["a"])(r,n,s,!1,null,"6b76b78c",null);t["default"]=l.exports}}]);