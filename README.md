# Generating events and patterns

_Work in progress_

This package aspires to generate
[Quamina](https://github.com/timbray/quamina)/[EventBridge
patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-patterns.html)
and events in a reasonably controllable way.

The basic idea is to specify the range of shapes for random events.
Then derive some patterns for those events so that matching sees
enough matches.

There's a command-line [tool](cmd/quaminagen) that tries to make it
easy to perform reasonably flexible, large-scale Quamina testing.

Here's some example output:

```
2022/05/21 14:50:24 made 1000 matching events
2022/05/21 14:50:24 made 1000 matching patterns
2022/05/21 14:50:24 made 1000 other events
2022/05/21 14:50:24 made 1000 other patterns
dims,op,atoms,bytes,atombytes,depth,width
data,op,atoms,bytes,atombytes,depth,width
dims,match,20,240,183,4,3
data,matchEvent,"{\"qxso\":{\"dqojdflctkwc\":-99,\"nmibmiqmw\":{\"bebzfogejpdihjyefmb\":[83,40,\"xvdnryyct\",-87],\"gxbedkkwtepriabq\":\"veibbfbvthyat\",\"ujsjudvxmh\":\"dqgkfhinupsmrseapkhbm\",\"wzpavgkupgczilyeh\":83}},\"wyvy\":{\"icktyxpxqvtvk\":-7,\"mxcfhizprrjkqhuytpb\":[-16]}}"
dims,match,5,93,71,3,2
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
dims,match,17,180,130,4,3
data,matchEvent,"{\"dgxlm\":59,\"dwcia\":[{\"lxxt\":-74,\"wgautdnk\":[\"uwjjyjagmkbgprklswq\",\"qfoqmlmycsudqryu\",\"kaxtfngiqqocdxxmbd\",40],\"wtjtdpdr\":[-45,\"yhtnymqndxcevnta\"],\"xmg\":\"rhvwsheelfcr\"}],\"zgt\":-36}"
dims,match,18,135,92,4,3
data,matchEvent,"{\"muglycdwtxaywulyw\":9,\"psdrl\":[\"dpsiqbrdntwjzpykxucygwe\",38,[[-55],[-24]],-19],\"xppwhiedwqabs\":[26,[-12,-29],[58,[-88,29,69],-26],35]}"
dims,match,5,93,71,3,2
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
dims,match,18,135,92,4,3
data,matchEvent,"{\"muglycdwtxaywulyw\":9,\"psdrl\":[\"dpsiqbrdntwjzpykxucygwe\",38,[[-55],[-24]],-19],\"xppwhiedwqabs\":[26,[-12,-29],[58,[-88,29,69],-26],35]}"
dims,match,17,180,130,4,3
data,matchEvent,"{\"dgxlm\":59,\"dwcia\":[{\"lxxt\":-74,\"wgautdnk\":[\"uwjjyjagmkbgprklswq\",\"qfoqmlmycsudqryu\",\"kaxtfngiqqocdxxmbd\",40],\"wtjtdpdr\":[-45,\"yhtnymqndxcevnta\"],\"xmg\":\"rhvwsheelfcr\"}],\"zgt\":-36}"
dims,add,5,115,93,3,2
data,addPattern,"{\"kcvhplqwgvkbkoxhmby\":{\"nqyskolejywoarytyw\":[\"fulwotpdkpegilwvz\"]},\"zjxyewldhs\":[\"tilpgwpikbcujewxlxddzyvltboty\"]}"
dims,match,47,675,517,6,5
data,matchEvent,"{\"lfwcugitf\":{\"rcyxxns\":\"jnphxthc\",\"sas\":[[-69,[\"mvqludtcszwpuxzifhwxksi\",[\"xzpzupfzwmgwyizrvehl\",\"kkuceuphhvmphwygbyrjen\"],-42],\"odbfbnbcfuzchlmvwj\"]],\"zcf\":\"vjeuyvvywhycvioolndgmdlkmos\"},\"mugjttaahem\":{\"altdglmk\":\"knxvmraynyimyeubtluyhaudn\",\"dtwpagbzvgfa\":[[[\"dgvuwjbbqbaczpeafzfha\"],[1,\"ddtnnpmqnjahdpokglwhep\"]],-77],\"rtouhvtsulrpw\":[[[93,\"flptvahfqcpseuvrivnso\"],\"spqfdictzehjppiilfnluqqjlvw\",\"yomdeemfwjyilerxuegnkrnajjca\"],[\"offbojsvthhpxfmnncmduvzeha\"],[-81,-13,[-86]]],\"vtqnlefthnayhinfdb\":-87},\"ygbvsaxwwuazjusf\":[48,[-79,[{\"unpbryphjze\":[\"mgethmgmhajueaxqthkuwp\",\"undgfowapzbxg\"],\"zcexn\":-23}],21,[11,[\"owubpugujjcbzsl\",[49],75],\"dezuc\",[\"tqlesganmxrevenv\",67]]]]}"
dims,match,47,675,517,6,5
dims,match,16,217,166,4,3
data,matchEvent,"{\"lfwcugitf\":{\"rcyxxns\":\"jnphxthc\",\"sas\":[[-69,[\"mvqludtcszwpuxzifhwxksi\",[\"xzpzupfzwmgwyizrvehl\",\"kkuceuphhvmphwygbyrjen\"],-42],\"odbfbnbcfuzchlmvwj\"]],\"zcf\":\"vjeuyvvywhycvioolndgmdlkmos\"},\"mugjttaahem\":{\"altdglmk\":\"knxvmraynyimyeubtluyhaudn\",\"dtwpagbzvgfa\":[[[\"dgvuwjbbqbaczpeafzfha\"],[1,\"ddtnnpmqnjahdpokglwhep\"]],-77],\"rtouhvtsulrpw\":[[[93,\"flptvahfqcpseuvrivnso\"],\"spqfdictzehjppiilfnluqqjlvw\",\"yomdeemfwjyilerxuegnkrnajjca\"],[\"offbojsvthhpxfmnncmduvzeha\"],[-81,-13,[-86]]],\"vtqnlefthnayhinfdb\":-87},\"ygbvsaxwwuazjusf\":[48,[-79,[{\"unpbryphjze\":[\"mgethmgmhajueaxqthkuwp\",\"undgfowapzbxg\"],\"zcexn\":-23}],21,[11,[\"owubpugujjcbzsl\",[49],75],\"dezuc\",[\"tqlesganmxrevenv\",67]]]]}"
dims,match,16,217,166,4,3
data,matchEvent,"{\"eadqtk\":-56,\"jkmlhdpmmic\":{\"lejhqibdmduqlnxnyt\":[\"vprdbutbzwmgflxo\",[38]],\"roghpd\":[\"ezhcdmzqfltl\",[91,-71,\"hlxcojhyjxrblg\"],{\"okpjllpwlhkix\":\"zfwlxxkwqkunnfzocbucualgqmf\",\"qkv\":31},\"hwcvtuaceltvpkisjzvdjdfablxt\"]}}"
dims,match,16,217,166,4,3
data,matchEvent,"{\"eadqtk\":-56,\"jkmlhdpmmic\":{\"lejhqibdmduqlnxnyt\":[\"vprdbutbzwmgflxo\",[38]],\"roghpd\":[\"ezhcdmzqfltl\",[91,-71,\"hlxcojhyjxrblg\"],{\"okpjllpwlhkix\":\"zfwlxxkwqkunnfzocbucualgqmf\",\"qkv\":31},\"hwcvtuaceltvpkisjzvdjdfablxt\"]}}"
dims,add,5,115,93,3,2
data,addPattern,"{\"kcvhplqwgvkbkoxhmby\":{\"nqyskolejywoarytyw\":[\"fulwotpdkpegilwvz\"]},\"zjxyewldhs\":[\"tilpgwpikbcujewxlxddzyvltboty\"]}"
data,matchEvent,"{\"eadqtk\":-56,\"jkmlhdpmmic\":{\"lejhqibdmduqlnxnyt\":[\"vprdbutbzwmgflxo\",[38]],\"roghpd\":[\"ezhcdmzqfltl\",[91,-71,\"hlxcojhyjxrblg\"],{\"okpjllpwlhkix\":\"zfwlxxkwqkunnfzocbucualgqmf\",\"qkv\":31},\"hwcvtuaceltvpkisjzvdjdfablxt\"]}}"
dims,match,5,93,71,3,2
dims,match,17,180,130,4,3
data,matchEvent,"{\"dgxlm\":59,\"dwcia\":[{\"lxxt\":-74,\"wgautdnk\":[\"uwjjyjagmkbgprklswq\",\"qfoqmlmycsudqryu\",\"kaxtfngiqqocdxxmbd\",40],\"wtjtdpdr\":[-45,\"yhtnymqndxcevnta\"],\"xmg\":\"rhvwsheelfcr\"}],\"zgt\":-36}"
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
dims,match,26,232,167,5,4
data,matchEvent,"{\"ndefanbwhqyg\":[-72,{\"nkmjcyxrktxo\":-6,\"xygr\":-42,\"zlvfqjhegyshmkykep\":[[-88,-13],-48,\"pncgqrgs\",\"snxga\"]},[73],\"uthvejizeslzeksyvdv\"],\"nyrdzropgeuepr\":[[80,{\"ficnquairvcmefui\":49,\"irsfclxkq\":-30,\"wlwldgujpxq\":-50},-24],49,-77,45]}"
dims,match,16,217,166,4,3
data,matchEvent,"{\"eadqtk\":-56,\"jkmlhdpmmic\":{\"lejhqibdmduqlnxnyt\":[\"vprdbutbzwmgflxo\",[38]],\"roghpd\":[\"ezhcdmzqfltl\",[91,-71,\"hlxcojhyjxrblg\"],{\"okpjllpwlhkix\":\"zfwlxxkwqkunnfzocbucualgqmf\",\"qkv\":31},\"hwcvtuaceltvpkisjzvdjdfablxt\"]}}"
dims,add,6,55,38,2,3
data,addPattern,"{\"bpfbda\":[-71],\"uyayide\":null,\"yzojsfaxgmpopttk\":[85]}"
dims,match,20,240,183,4,3
data,matchEvent,"{\"qxso\":{\"dqojdflctkwc\":-99,\"nmibmiqmw\":{\"bebzfogejpdihjyefmb\":[83,40,\"xvdnryyct\",-87],\"gxbedkkwtepriabq\":\"veibbfbvthyat\",\"ujsjudvxmh\":\"dqgkfhinupsmrseapkhbm\",\"wzpavgkupgczilyeh\":83}},\"wyvy\":{\"icktyxpxqvtvk\":-7,\"mxcfhizprrjkqhuytpb\":[-16]}}"
dims,add,6,55,38,2,3
data,addPattern,"{\"bpfbda\":[-71],\"uyayide\":null,\"yzojsfaxgmpopttk\":[85]}"
dims,add,5,115,93,3,2
data,addPattern,"{\"kcvhplqwgvkbkoxhmby\":{\"nqyskolejywoarytyw\":[\"fulwotpdkpegilwvz\"]},\"zjxyewldhs\":[\"tilpgwpikbcujewxlxddzyvltboty\"]}"
dims,match,18,135,92,4,3
data,matchEvent,"{\"muglycdwtxaywulyw\":9,\"psdrl\":[\"dpsiqbrdntwjzpykxucygwe\",38,[[-55],[-24]],-19],\"xppwhiedwqabs\":[26,[-12,-29],[58,[-88,29,69],-26],35]}"
dims,match,16,217,166,4,3
dims,match,5,93,71,3,2
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
data,matchEvent,"{\"eadqtk\":-56,\"jkmlhdpmmic\":{\"lejhqibdmduqlnxnyt\":[\"vprdbutbzwmgflxo\",[38]],\"roghpd\":[\"ezhcdmzqfltl\",[91,-71,\"hlxcojhyjxrblg\"],{\"okpjllpwlhkix\":\"zfwlxxkwqkunnfzocbucualgqmf\",\"qkv\":31},\"hwcvtuaceltvpkisjzvdjdfablxt\"]}}"
2022/05/21 14:50:24 1 executed 10 ops
dims,match,16,217,166,4,3
data,matchEvent,"{\"eadqtk\":-56,\"jkmlhdpmmic\":{\"lejhqibdmduqlnxnyt\":[\"vprdbutbzwmgflxo\",[38]],\"roghpd\":[\"ezhcdmzqfltl\",[91,-71,\"hlxcojhyjxrblg\"],{\"okpjllpwlhkix\":\"zfwlxxkwqkunnfzocbucualgqmf\",\"qkv\":31},\"hwcvtuaceltvpkisjzvdjdfablxt\"]}}"
dims,match,47,675,517,6,5
data,matchEvent,"{\"lfwcugitf\":{\"rcyxxns\":\"jnphxthc\",\"sas\":[[-69,[\"mvqludtcszwpuxzifhwxksi\",[\"xzpzupfzwmgwyizrvehl\",\"kkuceuphhvmphwygbyrjen\"],-42],\"odbfbnbcfuzchlmvwj\"]],\"zcf\":\"vjeuyvvywhycvioolndgmdlkmos\"},\"mugjttaahem\":{\"altdglmk\":\"knxvmraynyimyeubtluyhaudn\",\"dtwpagbzvgfa\":[[[\"dgvuwjbbqbaczpeafzfha\"],[1,\"ddtnnpmqnjahdpokglwhep\"]],-77],\"rtouhvtsulrpw\":[[[93,\"flptvahfqcpseuvrivnso\"],\"spqfdictzehjppiilfnluqqjlvw\",\"yomdeemfwjyilerxuegnkrnajjca\"],[\"offbojsvthhpxfmnncmduvzeha\"],[-81,-13,[-86]]],\"vtqnlefthnayhinfdb\":-87},\"ygbvsaxwwuazjusf\":[48,[-79,[{\"unpbryphjze\":[\"mgethmgmhajueaxqthkuwp\",\"undgfowapzbxg\"],\"zcexn\":-23}],21,[11,[\"owubpugujjcbzsl\",[49],75],\"dezuc\",[\"tqlesganmxrevenv\",67]]]]}"
dims,match,16,217,166,4,3
dims,match,5,93,71,3,2
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
2022/05/21 14:50:24 0 executed 10 ops
dims,add,5,115,93,3,2
data,addPattern,"{\"kcvhplqwgvkbkoxhmby\":{\"nqyskolejywoarytyw\":[\"fulwotpdkpegilwvz\"]},\"zjxyewldhs\":[\"tilpgwpikbcujewxlxddzyvltboty\"]}"
data,matchEvent,"{\"eadqtk\":-56,\"jkmlhdpmmic\":{\"lejhqibdmduqlnxnyt\":[\"vprdbutbzwmgflxo\",[38]],\"roghpd\":[\"ezhcdmzqfltl\",[91,-71,\"hlxcojhyjxrblg\"],{\"okpjllpwlhkix\":\"zfwlxxkwqkunnfzocbucualgqmf\",\"qkv\":31},\"hwcvtuaceltvpkisjzvdjdfablxt\"]}}"
dims,match,5,93,71,3,2
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
dims,match,47,675,517,6,5
data,matchEvent,"{\"lfwcugitf\":{\"rcyxxns\":\"jnphxthc\",\"sas\":[[-69,[\"mvqludtcszwpuxzifhwxksi\",[\"xzpzupfzwmgwyizrvehl\",\"kkuceuphhvmphwygbyrjen\"],-42],\"odbfbnbcfuzchlmvwj\"]],\"zcf\":\"vjeuyvvywhycvioolndgmdlkmos\"},\"mugjttaahem\":{\"altdglmk\":\"knxvmraynyimyeubtluyhaudn\",\"dtwpagbzvgfa\":[[[\"dgvuwjbbqbaczpeafzfha\"],[1,\"ddtnnpmqnjahdpokglwhep\"]],-77],\"rtouhvtsulrpw\":[[[93,\"flptvahfqcpseuvrivnso\"],\"spqfdictzehjppiilfnluqqjlvw\",\"yomdeemfwjyilerxuegnkrnajjca\"],[\"offbojsvthhpxfmnncmduvzeha\"],[-81,-13,[-86]]],\"vtqnlefthnayhinfdb\":-87},\"ygbvsaxwwuazjusf\":[48,[-79,[{\"unpbryphjze\":[\"mgethmgmhajueaxqthkuwp\",\"undgfowapzbxg\"],\"zcexn\":-23}],21,[11,[\"owubpugujjcbzsl\",[49],75],\"dezuc\",[\"tqlesganmxrevenv\",67]]]]}"
dims,match,17,180,130,4,3
data,matchEvent,"{\"dgxlm\":59,\"dwcia\":[{\"lxxt\":-74,\"wgautdnk\":[\"uwjjyjagmkbgprklswq\",\"qfoqmlmycsudqryu\",\"kaxtfngiqqocdxxmbd\",40],\"wtjtdpdr\":[-45,\"yhtnymqndxcevnta\"],\"xmg\":\"rhvwsheelfcr\"}],\"zgt\":-36}"
dims,match,5,93,71,3,2
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
2022/05/21 14:50:24 2 executed 10 ops
dims,match,18,135,92,4,3
data,matchEvent,"{\"muglycdwtxaywulyw\":9,\"psdrl\":[\"dpsiqbrdntwjzpykxucygwe\",38,[[-55],[-24]],-19],\"xppwhiedwqabs\":[26,[-12,-29],[58,[-88,29,69],-26],35]}"
dims,match,5,93,71,3,2
data,matchEvent,"{\"annbzcjwzg\":[[\"fnrlivyjcqfjasyng\"]],\"ztkjyln\":[\"qjdrjdyueyxhprtlizefaqdgsnmg\",\"eejpajgwf\"]}"
2022/05/21 14:50:24 3 executed 10 ops
{
  "Live": 4,
  "Added": 4,
  "Deleted": 0,
  "Emitted": 0,
  "Filtered": 0,
  "LastRebuilt": "0001-01-01T00:00:00Z",
  "RebuildDuration": 0,
  "RebuildPurged": 0
}
elapsed 2.154253ms
```
