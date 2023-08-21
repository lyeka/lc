
## config
### leetcode-editor
code filename
```
$!velocityTool.leftPadZeros(${question.frontendQuestionId},4)/$!velocityTool.leftPadZeros($!{question.frontendQuestionId},4)_$!velocityTool.smallCamelCaseName(${question.title})```
code template
```
package $!velocityTool.smallCamelCaseName(${question.titleSlug})

${question.code}
```

### auto generate question list
refer [gen_question_list.sh](./gen_question_list.sh)
