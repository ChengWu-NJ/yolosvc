### sample of config.yml
```yaml
darknetConfigFile: /root/go/src/github.com/zmisc/yolosvc/test/configgenerate/darknetfiles/nn.cfg
objNamesFile: /root/go/src/github.com/zmisc/yolosvc/test/configgenerate/darknetfiles/obj.names
darknetWeightsFile: /root/go/src/github.com/zmisc/yolosvc/test/configgenerate/darknetfiles/nn.weights
detectThreshold: 0.8
portOfGrpcSvc: 37658
objClasses:
- id: 0
  name: yellowGourami
  # yellow
  labelColorR: 255
  labelColorG: 255
  labelColorB: 0
- id: 1
  name: blueGourami
  # light blue
  labelColorR: 173
  labelColorG: 216
  labelColorB: 230
- id: 2
  name: redFighter
  # red
  labelColorR: 255
  labelColorG: 0
  labelColorB: 0
- id: 3
  name: originalFighter
  # orange
  labelColorR: 255
  labelColorG: 165
  labelColorB: 0
- id: 4
  name: fox
  # blue
  labelColorR: 0
  labelColorG: 0
  labelColorB: 255
- id: 5
  name: sucker
  # green
  labelColorR: 0
  labelColorG: 128
  labelColorB: 0
- id: 6
  name: tiger
  # dark red
  labelColorR: 139
  labelColorG: 0
  labelColorB: 0
- id: 7
  name: WCMM
  # pale green
  labelColorR: 152
  labelColorG: 251
  labelColorB: 152
```