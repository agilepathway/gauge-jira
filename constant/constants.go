/*----------------------------------------------------------------
 *  Copyright (c) ThoughtWorks, Inc.
 *  Licensed under the Apache License, Version 2.0
 *  See LICENSE in the project root for license information.
 *----------------------------------------------------------------*/
package constant

var CSS = `
body .specs ul ol {
    list-style-type: decimal !important;
}
/*# github markdown css from https://cdnjs.com/libraries/github-markdown-css */
@font-face{font-family:octicons-link;src:url(data:font/woff;charset=utf-8;base64,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) format('woff')}.markdown-body{-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;line-height:1.5;color:#24292e;font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol";font-size:16px;line-height:1.5;word-wrap:break-word}.markdown-body .pl-c{color:#6a737d}.markdown-body .pl-c1,.markdown-body .pl-s .pl-v{color:#005cc5}.markdown-body .pl-e,.markdown-body .pl-en{color:#6f42c1}.markdown-body .pl-s .pl-s1,.markdown-body .pl-smi{color:#24292e}.markdown-body .pl-ent{color:#22863a}.markdown-body .pl-k{color:#d73a49}.markdown-body .pl-pds,.markdown-body .pl-s,.markdown-body .pl-s .pl-pse .pl-s1,.markdown-body .pl-sr,.markdown-body .pl-sr .pl-cce,.markdown-body .pl-sr .pl-sra,.markdown-body .pl-sr .pl-sre{color:#032f62}.markdown-body .pl-smw,.markdown-body .pl-v{color:#e36209}.markdown-body .pl-bu{color:#b31d28}.markdown-body .pl-ii{color:#fafbfc;background-color:#b31d28}.markdown-body .pl-c2{color:#fafbfc;background-color:#d73a49}.markdown-body .pl-c2::before{content:"^M"}.markdown-body .pl-sr .pl-cce{font-weight:700;color:#22863a}.markdown-body .pl-ml{color:#735c0f}.markdown-body .pl-mh,.markdown-body .pl-mh .pl-en,.markdown-body .pl-ms{font-weight:700;color:#005cc5}.markdown-body .pl-mi{font-style:italic;color:#24292e}.markdown-body .pl-mb{font-weight:700;color:#24292e}.markdown-body .pl-md{color:#b31d28;background-color:#ffeef0}.markdown-body .pl-mi1{color:#22863a;background-color:#f0fff4}.markdown-body .pl-mc{color:#e36209;background-color:#ffebda}.markdown-body .pl-mi2{color:#f6f8fa;background-color:#005cc5}.markdown-body .pl-mdr{font-weight:700;color:#6f42c1}.markdown-body .pl-ba{color:#586069}.markdown-body .pl-sg{color:#959da5}.markdown-body .pl-corl{text-decoration:underline;color:#032f62}.markdown-body .octicon{display:inline-block;vertical-align:text-top;fill:currentColor}.markdown-body a{background-color:transparent;-webkit-text-decoration-skip:objects}.markdown-body a:active,.markdown-body a:hover{outline-width:0}.markdown-body strong{font-weight:inherit}.markdown-body strong{font-weight:bolder}.markdown-body h1{font-size:2em;margin:.67em 0}.markdown-body img{border-style:none}.markdown-body svg:not(:root){overflow:hidden}.markdown-body code,.markdown-body kbd,.markdown-body pre{font-family:monospace,monospace;font-size:1em}.markdown-body hr{box-sizing:content-box;height:0;overflow:visible}.markdown-body input{font:inherit;margin:0}.markdown-body input{overflow:visible}.markdown-body [type=checkbox]{box-sizing:border-box;padding:0}.markdown-body *{box-sizing:border-box}.markdown-body input{font-family:inherit;font-size:inherit;line-height:inherit}.markdown-body a{color:#0366d6;text-decoration:none}.markdown-body a:hover{text-decoration:underline}.markdown-body strong{font-weight:600}.markdown-body hr{height:0;margin:15px 0;overflow:hidden;background:0 0;border:0;border-bottom:1px solid #dfe2e5}.markdown-body hr::before{display:table;content:""}.markdown-body hr::after{display:table;clear:both;content:""}.markdown-body table{border-spacing:0;border-collapse:collapse}.markdown-body td,.markdown-body th{padding:0}.markdown-body h1,.markdown-body h2,.markdown-body h3,.markdown-body h4,.markdown-body h5,.markdown-body h6{margin-top:0;margin-bottom:0}.markdown-body h1{font-size:32px;font-weight:600}.markdown-body h2{font-size:24px;font-weight:600}.markdown-body h3{font-size:20px;font-weight:600}.markdown-body h4{font-size:16px;font-weight:600}.markdown-body h5{font-size:14px;font-weight:600}.markdown-body h6{font-size:12px;font-weight:600}.markdown-body p{margin-top:0;margin-bottom:10px}.markdown-body blockquote{margin:0}.markdown-body ol,.markdown-body ul{padding-left:0;margin-top:0;margin-bottom:0}.markdown-body ol ol,.markdown-body ul ol{list-style-type:lower-roman}.markdown-body ol ol ol,.markdown-body ol ul ol,.markdown-body ul ol ol,.markdown-body ul ul ol{list-style-type:lower-alpha}.markdown-body dd{margin-left:0}.markdown-body code{font-family:SFMono-Regular,Consolas,"Liberation Mono",Menlo,Courier,monospace;font-size:12px}.markdown-body pre{margin-top:0;margin-bottom:0;font:12px SFMono-Regular,Consolas,"Liberation Mono",Menlo,Courier,monospace}.markdown-body .octicon{vertical-align:text-bottom}.markdown-body .pl-0{padding-left:0!important}.markdown-body .pl-1{padding-left:4px!important}.markdown-body .pl-2{padding-left:8px!important}.markdown-body .pl-3{padding-left:16px!important}.markdown-body .pl-4{padding-left:24px!important}.markdown-body .pl-5{padding-left:32px!important}.markdown-body .pl-6{padding-left:40px!important}.markdown-body::before{display:table;content:""}.markdown-body::after{display:table;clear:both;content:""}.markdown-body>:first-child{margin-top:0!important}.markdown-body>:last-child{margin-bottom:0!important}.markdown-body a:not([href]){color:inherit;text-decoration:none}.markdown-body .anchor{float:left;padding-right:4px;margin-left:-20px;line-height:1}.markdown-body .anchor:focus{outline:0}.markdown-body blockquote,.markdown-body dl,.markdown-body ol,.markdown-body p,.markdown-body pre,.markdown-body table,.markdown-body ul{margin-top:0;margin-bottom:16px}.markdown-body hr{height:.25em;padding:0;margin:24px 0;background-color:#e1e4e8;border:0}.markdown-body blockquote{padding:0 1em;color:#6a737d;border-left:.25em solid #dfe2e5}.markdown-body blockquote>:first-child{margin-top:0}.markdown-body blockquote>:last-child{margin-bottom:0}.markdown-body kbd{display:inline-block;padding:3px 5px;font-size:11px;line-height:10px;color:#444d56;vertical-align:middle;background-color:#fafbfc;border:solid 1px #c6cbd1;border-bottom-color:#959da5;border-radius:3px;box-shadow:inset 0 -1px 0 #959da5}.markdown-body h1,.markdown-body h2,.markdown-body h3,.markdown-body h4,.markdown-body h5,.markdown-body h6{margin-top:24px;margin-bottom:16px;font-weight:600;line-height:1.25}.markdown-body h1 .octicon-link,.markdown-body h2 .octicon-link,.markdown-body h3 .octicon-link,.markdown-body h4 .octicon-link,.markdown-body h5 .octicon-link,.markdown-body h6 .octicon-link{color:#1b1f23;vertical-align:middle;visibility:hidden}.markdown-body h1:hover .anchor,.markdown-body h2:hover .anchor,.markdown-body h3:hover .anchor,.markdown-body h4:hover .anchor,.markdown-body h5:hover .anchor,.markdown-body h6:hover .anchor{text-decoration:none}.markdown-body h1:hover .anchor .octicon-link,.markdown-body h2:hover .anchor .octicon-link,.markdown-body h3:hover .anchor .octicon-link,.markdown-body h4:hover .anchor .octicon-link,.markdown-body h5:hover .anchor .octicon-link,.markdown-body h6:hover .anchor .octicon-link{visibility:visible}.markdown-body h1{padding-bottom:.3em;font-size:2em;border-bottom:1px solid #eaecef}.markdown-body h2{padding-bottom:.3em;font-size:1.5em;border-bottom:1px solid #eaecef}.markdown-body h3{font-size:1.25em}.markdown-body h4{font-size:1em}.markdown-body h5{font-size:.875em}.markdown-body h6{font-size:.85em;color:#6a737d}.markdown-body ol,.markdown-body ul{padding-left:2em}.markdown-body ol ol,.markdown-body ol ul,.markdown-body ul ol,.markdown-body ul ul{margin-top:0;margin-bottom:0}.markdown-body li>p{margin-top:16px}.markdown-body li+li{margin-top:.25em}.markdown-body dl{padding:0}.markdown-body dl dt{padding:0;margin-top:16px;font-size:1em;font-style:italic;font-weight:600}.markdown-body dl dd{padding:0 16px;margin-bottom:16px}.markdown-body table{display:block;width:100%;overflow:auto}.markdown-body table th{font-weight:600}.markdown-body table td,.markdown-body table th{padding:6px 13px;border:1px solid #dfe2e5}.markdown-body table tr{background-color:#fff;border-top:1px solid #c6cbd1}.markdown-body table tr:nth-child(2n){background-color:#f6f8fa}.markdown-body img{max-width:100%;box-sizing:content-box;background-color:#fff}.markdown-body code{padding:0;padding-top:.2em;padding-bottom:.2em;margin:0;font-size:85%;background-color:rgba(27,31,35,.05);border-radius:3px}.markdown-body code::after,.markdown-body code::before{letter-spacing:-.2em;content:"\00a0"}.markdown-body pre{word-wrap:normal}.markdown-body pre>code{padding:0;margin:0;font-size:100%;word-break:normal;white-space:pre;background:0 0;border:0}.markdown-body .highlight{margin-bottom:16px}.markdown-body .highlight pre{margin-bottom:0;word-break:normal}.markdown-body .highlight pre,.markdown-body pre{padding:16px;overflow:auto;font-size:85%;line-height:1.45;background-color:#f6f8fa;border-radius:3px}.markdown-body pre code{display:inline;max-width:auto;padding:0;margin:0;overflow:visible;line-height:inherit;word-wrap:normal;background-color:transparent;border:0}.markdown-body pre code::after,.markdown-body pre code::before{content:normal}.markdown-body .full-commit .btn-outline:not(:disabled):hover{color:#005cc5;border-color:#005cc5}.markdown-body kbd{display:inline-block;padding:3px 5px;font:11px SFMono-Regular,Consolas,"Liberation Mono",Menlo,Courier,monospace;line-height:10px;color:#444d56;vertical-align:middle;background-color:#fafbfc;border:solid 1px #d1d5da;border-bottom-color:#c6cbd1;border-radius:3px;box-shadow:inset 0 -1px 0 #c6cbd1}.markdown-body :checked+.radio-label{position:relative;z-index:1;border-color:#0366d6}.markdown-body .task-list-item{list-style-type:none}.markdown-body .task-list-item+.task-list-item{margin-top:3px}.markdown-body .task-list-item input{margin:0 .2em .25em -1.6em;vertical-align:middle}.markdown-body hr{border-bottom-color:#eee}
/*# sourceMappingURL=github-markdown.min.css.map */
.markdown-body {
    min-width: 200px;
    max-width: 980px;
    margin: 0 auto;
    padding: 45px;
    padding-top: 20px;
}
@media (max-width: 767px) {
	.markdown-body {
		padding: 15px;
	}
}
ul#navigation li a {
	display: block;
	width: 5em;
	color: #FFF;
	background-color: #666666;
	padding: .2em 0;
	text-align: center;
	text-decoration: none;
}

ul#navigation li a:hover {
	color: #FFF;
	background-color: #878686;
}
.tags {
	border: 1px solid #848484;
    -webkit-border-radius: 30px;
    -moz-border-radius: 30px;
    border-radius: 10px;
    outline: 0;
    height: 37px;
    width: 100%;
    padding-left: 10px;
    padding-right: 10px;
}
ul#navigation .nav {
	display: inline-block;
	margin: 5px;
}
.entity {
	font-weight: 500;
}
.stats {
	font-style: italic;
	font-size: 18px;
    text-align: center;
    background-color: #f3f3f3;
    border-radius: 4px;
    margin-top: 2% !important;
    padding-bottom: 4px;
    padding-top: 4px;
}
.separator {
	background: black;
	margin-left: 1%;
    margin-right: 0.3%;
}
.specs p {
	margin: 0px;
}
input.tags::placeholder {
    font-size: 13px;
    font-style: italic;
}
h1 {
	margin-top: 5px !important;
}
`
var IndexJS = `<script src="index.js"></script>`

var IndexJSContent = `
const statsTemplate = '<p class="stats">\
<span>\
	<span class="entity">Specifications: </span>\
	<span class="stat">SPEC_NUMBER</span>\
</span>\
<span class="separator">|</span>\
<span>\
	<span class="entity">Scenarios: </span>\
	<span class="stat">SCENARIO_NUMBER</span>\
</span>\
</p>\
';
const populateIndex = (specs) => {
    if (Object.keys(specs).length == 0) {
		document.getElementsByClassName("specs")[0].innerHTML = statsTemplate.replace("SPEC_NUMBER", 0).replace("SCENARIO_NUMBER", 0);
        return;
    }
    let text = "<ul>";
    let scenarioNumber = 0;
    specs.forEach((spec) => {
        text += "<li><p><b><a href=\"" + spec["path"] + "\">" + spec["name"] + "</a></b><ol>";
        spec.scenarios.forEach((scn) => {
            text += "<li>" + scn["name"] + "</li>";
            scenarioNumber++;
        });
        text += "</ol></p></li>";
    });
    text += "</ul></div>";
    const stats = statsTemplate.replace("SPEC_NUMBER", specs.length).replace("SCENARIO_NUMBER", scenarioNumber);
    document.getElementsByClassName("specs")[0].innerHTML = stats + text;
}

function handle(e) {
    if (e.keyCode === 13) {
        if (document.getElementsByClassName("tags")[0].value.trim() === "") {
            populateIndex(specs);
            return false;
        }
        populateIndex(filterSpecs(document.getElementsByClassName("tags")[0].value));
    }
    return false;
}
const filterSpecs = (tagExp) => {
    tags = getTagsWithoutOperators(tagExp).map((e) => e.replace("!", ""));
    const newSpecs = [];
    specs.forEach((spec) => {
        const scenarios = [];
        spec.scenarios.forEach(function(scn) {
            let newTagExp = replace(tagExp, scn.tags.filter(function(t) {
                return t !== "";
            }), "true");
            newTagExp = replace(newTagExp, tags, "false");
            if (eval(newTagExp)) scenarios.push(scn);
        });
        if (scenarios.length > 0)
            newSpecs.push({ "name": spec.name, "path": spec.path, "scenarios": scenarios });
    });
    return newSpecs;
}
const replace = (tagExp, tags, replaceString) => {
    let tagsWithOperators = getTagsWithOperators(tagExp);
    tags.forEach(function(t) {
        tagsWithOperators = replaceAll(tagsWithOperators, t, replaceString)
        tagsWithOperators = replaceAll(tagsWithOperators, "!" + t, "!" + replaceString)
    });
    return tagsWithOperators.join("");
}
const replaceAll = (elements, element, replacement) => elements.map((e) => element == e ? replacement : e);
const getTags = (tagExp, regex) => tagExp.split(regex).map((e) => e.trim()).filter((t) => t !== "");
const getTagsWithOperators = (tagExp) => getTags(tagExp, /(\)|\(|&|\|)/);
const getTagsWithoutOperators = (tagExp) => getTags(tagExp.replace(/\(/g, "").replace(/\)/g, ""), /(?:&|\|)/);
populateIndex(specs);
`

var IndexContent = `<center>
<h1>%s</h1>
	<input type="text" class="tags" placeholder="Tag expression to Filter specs/scenarios. Example: 'login & product | !customer'" onkeypress="handle(event)"></input>
</center>
<div class="specs"></div>
`
var DataFile = `<script src="data.js"></script>`
