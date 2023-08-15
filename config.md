
## config
### leetcode-editor
code filename
```
${question.frontendQuestionId}/$!{question.frontendQuestionId}_$!velocityTool.smallCamelCaseName(${question.title})
```
code template
```
package $!velocityTool.smallCamelCaseName(${question.titleSlug})

${question.code}
```

### auto generate question list
refer [gen_question_list.sh](./gen_question_list.sh)
