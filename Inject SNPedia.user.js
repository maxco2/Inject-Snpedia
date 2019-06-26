// ==UserScript==
// @name         Inject SNPedia
// @namespace    http://stateoftheart.pw/
// @version      0.1.0
// @description  inject snpedia
// @description: zh-CN inject snpedia
// @author       maxco2
// @match        *://www.snpedia.com/index.php/*
// @grant        GM_xmlhttpRequest
// @connect      127.0.0.1
// @connect      www.snpedia.com
// ==/UserScript==

function getURL_GM(url, callback) {
    GM_xmlhttpRequest({
        method: 'GET',
        url: url,
        onload: function(response) {
            if (response.status >= 200 && response.status < 400){
                callback(response.responseText);
            }
            else{
                console.log('Error getting ' + url + ' (' + this.status + ' ' + this.statusText + '): ' + this.responseText);
            }
        },
        onerror: function(response) {
            console.log('Error during GM_xmlhttpRequest to ' + url + ': ' + response.statusText);
        }
    });
}

function getJSON_GM(url, callback) {
    getURL_GM(url, function(data) {
        callback(JSON.parse(data));
    });
}


(function () {
    var host = location.hostname;
    if (host === "www.snpedia.com") {
        var rsidRe = location.href.match(/[rR][sS]\d+/);
        if (rsidRe) {
            var rsid=rsidRe[0].toLowerCase();
            getJSON_GM("http://127.0.0.1:8080/api/rsid/"+rsid,function(data){
                console.log('try inject snpedia');
                var str="";
                console.log(data);
                if(Array.isArray(data))
                {
                    for(var i=0;i<data.length;++i)
                    {
                        str+="genotype:";
                        str+=data[i].genotype;
                        str+=" ";
                    }
                }
                else
                {
                    str+="genotype:";
                    str+=data.genotype;
                }
                console.log(str);
                document.getElementById("firstHeading").innerHTML=rsid+" "+str;
            })
        }
    }
})();
