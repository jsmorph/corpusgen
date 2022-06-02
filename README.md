# Generating events and patterns

_Work in progress_

This package aspires to generate
[Quamina](https://github.com/timbray/quamina) and quasi-[EventBridge
patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html)
and events in a reasonably controllable way.  (Note that EventBridge
patterns have some required structure, and this generator currently
doesn't know about that.)

The basic idea is to specify the range of shapes for values in random
events.  Then derive some patterns for those events so that matching
sees enough matches.

There's a command-line [tool](cmd/corpusgen) that writes generated
events and patterns to its `stdout`.  See `corpusgen -h` for terse
documentation:

```Shell
Usage of corpusgen:
  -cfg
    	print cfg and exit
  -cfgfile string
    	filename for cfg JOSN
  -cfgjson string
    	JSON cfg overlay (default "{}")
  -dims
    	emit dimensions
  -events int
    	number of events (default 10)
  -nonmaps
    	events and patterns don't have to be maps
  -pats-per-event int
    	number of patterns per event (default 2)
  -patterns int
    	number of independent patterns (default 5)
  -seed int
    	random number generator seed (default 1654207335451164943)
```


For example:

```Shell
corpusgen -seed 42 -dims -cfgjson '{"Value":{"Maps":{"Properties":{"Cardinality":10}}}}'
```

will probably produce

```
event {"VURsMMNZJVSpaUYCO":{"kpTTueZPt":56,"tJlMjdrYzYIShBHY":{"KfBSWYjUNGkdmdt":{"VURsMMNZJVSpaUYCO":34,"gghEjkMVeZJpmKqakm":"ivDXQBaqLUjOW","tJlMjdrYzYIShBHY":[[60,35,-4,[-23,"AdEDKaxFyKzoVIquw",["SmOXPDnqnxWzUTO",15,"qjZOWnYx"],"uptjZxlSCpUKdSo"]],[[64,"MKZJmvs"],"FrnalxnyToTagOJtDWnrqJ"],"QqxWFljQIQOqzWQLQiuKlAE"]},"kpTTueZPt":-40},"wBK":68},"kpTTueZPt":[-47],"nhuksqVGzA":{"HadsNpb":46,"KfBSWYjUNGkdmdt":["wRQZW"],"gghEjkMVeZJpmKqakm":19},"tJlMjdrYzYIShBHY":82}
event dims {"Atoms":38,"Bytes":462,"AtomBytes":351,"Depth":8,"Width":7}
pattern {"VURsMMNZJVSpaUYCO":{"kpTTueZPt":[56],"wBK":[68]},"kpTTueZPt":[-47]}
pattern dims {"Atoms":38,"Bytes":462,"AtomBytes":351,"Depth":8,"Width":7}
pattern {"kpTTueZPt":[-47],"tJlMjdrYzYIShBHY":[82]}
pattern dims {"Atoms":38,"Bytes":462,"AtomBytes":351,"Depth":8,"Width":7}
event {"KfBSWYjUNGkdmdt":40,"kunrYaXseiLVpPCjZ":{"HadsNpb":87,"tJlMjdrYzYIShBHY":[-41],"wBK":["HdULZMyawiqCQmNXPmhdrtz",["JoJkOtnnUjoIhnkkci"]]}}
event dims {"Atoms":10,"Bytes":139,"AtomBytes":106,"Depth":4,"Width":3}
pattern {"kunrYaXseiLVpPCjZ":{"HadsNpb":[87]}}
pattern dims {"Atoms":10,"Bytes":139,"AtomBytes":106,"Depth":4,"Width":3}
pattern {"KfBSWYjUNGkdmdt":[40]}
pattern dims {"Atoms":10,"Bytes":139,"AtomBytes":106,"Depth":4,"Width":3}
event {"DsDAZpMp":-63,"HadsNpb":{"VURsMMNZJVSpaUYCO":-20,"nhuksqVGzA":94}}
event dims {"Atoms":7,"Bytes":68,"AtomBytes":50,"Depth":2,"Width":2}
pattern {"DsDAZpMp":[-63]}
pattern dims {"Atoms":7,"Bytes":68,"AtomBytes":50,"Depth":2,"Width":2}
pattern {"DsDAZpMp":[-63]}
pattern dims {"Atoms":7,"Bytes":68,"AtomBytes":50,"Depth":2,"Width":2}
event {"gghEjkMVeZJpmKqakm":[{"VURsMMNZJVSpaUYCO":[["pLCuAtsLxRyiRzDnY",44,48],39,"lpYICBiYkOJoXgGcLLLIGzKUokX"],"kpTTueZPt":"LBiXmkg","kunrYaXseiLVpPCjZ":["FvEaEGMQBTybCMzIjsuOcb",-33,"kUSzFMoCVrVfUKHqVaH","fVxsXalFkkrKSKiNnkLNWHnITl"],"wBK":21},["fcfMAmBjgMAqfcI"]],"nhuksqVGzA":{"HadsNpb":[-70,49,-78],"VURsMMNZJVSpaUYCO":[80,[[["kEEdEjlxMrkOVLNVcgWL","CnqGCtRZwcfFsLeRDZQElesOva",-13],-15,82]],0]}}
event dims {"Atoms":30,"Bytes":396,"AtomBytes":307,"Depth":6,"Width":5}
pattern {"nhuksqVGzA":{"HadsNpb":[-78],"VURsMMNZJVSpaUYCO":[0,80]}}
pattern dims {"Atoms":30,"Bytes":396,"AtomBytes":307,"Depth":6,"Width":5}
pattern {"nhuksqVGzA":{"VURsMMNZJVSpaUYCO":[80]}}
pattern dims {"Atoms":30,"Bytes":396,"AtomBytes":307,"Depth":6,"Width":5}
event {"KfBSWYjUNGkdmdt":{"KfBSWYjUNGkdmdt":[[84,"zpFKgArdSgUxIYLoXhIh","wkXankNiDnMmNYqTt"],"IWEHHxsEwZMQyB",["DqQDmuuORqaMQyJhSOgOWiMgAIDM",["mRcQNhzwwZHkfUzcWRyRcImvrYnqU","eVSWqi"]]],"gghEjkMVeZJpmKqakm":65,"wBK":[45]},"nhuksqVGzA":-4,"tJlMjdrYzYIShBHY":-51}
event dims {"Atoms":17,"Bytes":256,"AtomBytes":202,"Depth":5,"Width":4}
pattern {"KfBSWYjUNGkdmdt":{"gghEjkMVeZJpmKqakm":[65]},"tJlMjdrYzYIShBHY":[-51]}
pattern dims {"Atoms":17,"Bytes":256,"AtomBytes":202,"Depth":5,"Width":4}
pattern {"KfBSWYjUNGkdmdt":{"KfBSWYjUNGkdmdt":["IWEHHxsEwZMQyB"],"gghEjkMVeZJpmKqakm":[65],"wBK":null},"nhuksqVGzA":[-4],"tJlMjdrYzYIShBHY":[-51]}
pattern dims {"Atoms":17,"Bytes":256,"AtomBytes":202,"Depth":5,"Width":4}
event {"kunrYaXseiLVpPCjZ":[47,["tcQQiXRJqqmMNYLjCbjMdQe",["srSsXUU"]],["raNaZDdax"]],"nhuksqVGzA":{"KfBSWYjUNGkdmdt":{"DsDAZpMp":[-13,"JlIICBbdTHiChofncwDujEA"],"gghEjkMVeZJpmKqakm":-58,"wBK":-31},"VURsMMNZJVSpaUYCO":{"HadsNpb":"EDtYhVrZdqHeFBBpHWAi","VURsMMNZJVSpaUYCO":-67,"kunrYaXseiLVpPCjZ":[-6],"tJlMjdrYzYIShBHY":{"HadsNpb":"IWfJhFmxZDxFXsBTWIhaSueYDREov","kpTTueZPt":"SvCaDaeAETSsntumwlnYtsNLwX","wBK":"qdCTbwKs"}},"gghEjkMVeZJpmKqakm":[-88,-36,"AigRYQNZxIevkhHdS"],"kpTTueZPt":[67,"WoOiRDerMiZJlWvdAB",-71]},"tJlMjdrYzYIShBHY":{"DsDAZpMp":"LslvsVxDbkhMH","KfBSWYjUNGkdmdt":41,"VURsMMNZJVSpaUYCO":-85}}
event dims {"Atoms":43,"Bytes":604,"AtomBytes":472,"Depth":4,"Width":3}
pattern {"kunrYaXseiLVpPCjZ":null,"nhuksqVGzA":{"KfBSWYjUNGkdmdt":{"DsDAZpMp":[-13]},"VURsMMNZJVSpaUYCO":{"kunrYaXseiLVpPCjZ":[-6],"tJlMjdrYzYIShBHY":{"HadsNpb":["IWfJhFmxZDxFXsBTWIhaSueYDREov"]}},"gghEjkMVeZJpmKqakm":[-36,"AigRYQNZxIevkhHdS"]}}
pattern dims {"Atoms":43,"Bytes":604,"AtomBytes":472,"Depth":4,"Width":3}
pattern {"kunrYaXseiLVpPCjZ":null,"nhuksqVGzA":{"gghEjkMVeZJpmKqakm":[-36,-88]},"tJlMjdrYzYIShBHY":{"DsDAZpMp":["LslvsVxDbkhMH"]}}
pattern dims {"Atoms":43,"Bytes":604,"AtomBytes":472,"Depth":4,"Width":3}
event {"DsDAZpMp":[-60,["rkcEOhsvZUjwsTWIeLxiElYgwD",[["tIvMZehSGqJnrputvukxZlA",63,21,["qlsRkFwCO",-27,"bFuOduWQjOIILbvqSdvzoHfLh",93]]]],55],"nhuksqVGzA":{"KfBSWYjUNGkdmdt":82,"kunrYaXseiLVpPCjZ":75,"nhuksqVGzA":[-82],"tJlMjdrYzYIShBHY":"fhyqyKLTa"}}
event dims {"Atoms":20,"Bytes":246,"AtomBytes":189,"Depth":6,"Width":5}
pattern {"DsDAZpMp":null}
pattern dims {"Atoms":20,"Bytes":246,"AtomBytes":189,"Depth":6,"Width":5}
pattern {"DsDAZpMp":[55]}
pattern dims {"Atoms":20,"Bytes":246,"AtomBytes":189,"Depth":6,"Width":5}
event {"DsDAZpMp":-72,"KfBSWYjUNGkdmdt":{"KfBSWYjUNGkdmdt":{"DsDAZpMp":79,"KfBSWYjUNGkdmdt":"hhkxqBMJWYXFFOXYPEU","nhuksqVGzA":"QUZhAQFOKoqPkUJTOTusPZhdwgp","wBK":-53},"VURsMMNZJVSpaUYCO":{"KfBSWYjUNGkdmdt":87,"VURsMMNZJVSpaUYCO":95,"kunrYaXseiLVpPCjZ":["zgeVkFQDFFT"]},"wBK":23},"nhuksqVGzA":[43,14,37]}
event dims {"Atoms":25,"Bytes":298,"AtomBytes":230,"Depth":4,"Width":3}
pattern {"DsDAZpMp":[-72],"KfBSWYjUNGkdmdt":{"VURsMMNZJVSpaUYCO":{"KfBSWYjUNGkdmdt":[87],"VURsMMNZJVSpaUYCO":[95],"kunrYaXseiLVpPCjZ":null}},"nhuksqVGzA":[37]}
pattern dims {"Atoms":25,"Bytes":298,"AtomBytes":230,"Depth":4,"Width":3}
pattern {"DsDAZpMp":[-72],"KfBSWYjUNGkdmdt":{"wBK":[23]}}
pattern dims {"Atoms":25,"Bytes":298,"AtomBytes":230,"Depth":4,"Width":3}
event {"KfBSWYjUNGkdmdt":[-94],"tJlMjdrYzYIShBHY":{"DsDAZpMp":{"DsDAZpMp":41,"VURsMMNZJVSpaUYCO":["vUufjvcbNKBRj",{"DsDAZpMp":"FzsRPiYKRbVYT","VURsMMNZJVSpaUYCO":7,"gghEjkMVeZJpmKqakm":-36,"wBK":[["OKMcbvibrmnTlpGktksqbaYvwI",[[71,"ohotUnksrqXpERkGTEsJDOTePwUTf","oQmPLPnHVSdV"],"EfXUlhqs","HbAonrijkLFAwTSgaPNfyRshUnpXN","jMUKUxfIPbxwnU"],23,"nnXUKghkxodmbZjWUt"],"aYLyp"]},["iHpbEqTablDwjf"],"TYnvdMrOd"],"nhuksqVGzA":["iXpbBwgz",{"kpTTueZPt":"qJSjHrfaGgDUFkq","tJlMjdrYzYIShBHY":"prpjYBaeIqxKIWxIqzMHTXQ"},[["zeVFtxlcUuLtyTVAaQlMXYMEXi",["ZvbtVfbIIxpRNLWtMyGh","GWIQMeMYromc","MATLLJN",-93],-89]]]},"HadsNpb":["oxxdRViSKTA"],"wBK":["gTPWndZFtZk",-28,19,-75]}}
event dims {"Atoms":46,"Bytes":656,"AtomBytes":505,"Depth":9,"Width":8}
pattern {"tJlMjdrYzYIShBHY":{"DsDAZpMp":{"nhuksqVGzA":["iXpbBwgz"]},"HadsNpb":null}}
pattern dims {"Atoms":46,"Bytes":656,"AtomBytes":505,"Depth":9,"Width":8}
pattern {"KfBSWYjUNGkdmdt":null}
pattern dims {"Atoms":46,"Bytes":656,"AtomBytes":505,"Depth":9,"Width":8}
event {"gghEjkMVeZJpmKqakm":{"HadsNpb":62,"VURsMMNZJVSpaUYCO":-27},"kpTTueZPt":-88}
event dims {"Atoms":7,"Bytes":77,"AtomBytes":59,"Depth":2,"Width":2}
pattern {"gghEjkMVeZJpmKqakm":{"VURsMMNZJVSpaUYCO":[-27]}}
pattern dims {"Atoms":7,"Bytes":77,"AtomBytes":59,"Depth":2,"Width":2}
pattern {"gghEjkMVeZJpmKqakm":{"VURsMMNZJVSpaUYCO":[-27]}}
pattern dims {"Atoms":7,"Bytes":77,"AtomBytes":59,"Depth":2,"Width":2}
pattern {"kpTTueZPt":{"DsDAZpMp":["DoKstuITHNVF"],"HadsNpb":{"VURsMMNZJVSpaUYCO":["FvqURfObMXHYgeABAWIs"]},"nhuksqVGzA":{"VURsMMNZJVSpaUYCO":[-87]}}}
pattern dims {"Atoms":61,"Bytes":726,"AtomBytes":548,"Depth":5,"Width":4}
pattern {"gghEjkMVeZJpmKqakm":null,"wBK":null}
pattern dims {"Atoms":55,"Bytes":744,"AtomBytes":572,"Depth":7,"Width":6}
pattern {"nhuksqVGzA":[["YqvfFezeBcB"]]}
pattern dims {"Atoms":7,"Bytes":81,"AtomBytes":61,"Depth":3,"Width":3}
pattern {"kunrYaXseiLVpPCjZ":{"VURsMMNZJVSpaUYCO":[-13],"nhuksqVGzA":null}}
pattern dims {"Atoms":73,"Bytes":874,"AtomBytes":662,"Depth":6,"Width":5}
pattern {"wBK":[84]}
pattern dims {"Atoms":4,"Bytes":26,"AtomBytes":17,"Depth":1,"Width":2}
```
