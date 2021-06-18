# Ozon Code Platform Presentation Service

## Feedback GRPC API

#### CreateFeedbackV1
Позволяет создать новый фидбек.
##### Request
Содержит структуру типа ```Feedback```, обладающую полями:  
```Id``` - идентификатор фидбека, не используется в запросе    
```UserId``` - идентификатор пользователя, оставившего фидбек  
```ClassroomId``` - идентификатор классной комнаты  
```CommentId``` - комментарий, содержащий фидбек  
##### Response  
```FeedbackId``` - идентификатор созданного фидбека  

#### CreateMultiFeedbackV1  
Позволяет произвести пакетное создание фидбеков  
##### Request  
Содержит 1..N структур типа ```Feedback```  
##### Response  
Содержит следующие по порядку идентификаторы созданных фидбеков.

#### RemoveFeedbackV1
Позволяет удалить фидбек
##### Request
```FeedbackId``` - идентификатор фидбека, подлежащего удалению
##### Response
При удалении несуществующего фидбека будет возвращен соответствующий код GRPC 

#### DescribeFeedbackV1
Позволяет получить информацию о фидбеке
##### Request
```FeedbackId``` - идентификатор фидбека
##### Response
В случае отсутствия фидбека будет возвращен соответствующий код GRPC  
Содержит структуру типа ```Feedback```


#### UpdateFeedbackV1
Позволяет обновить информацию о фидбеке. При выполнении данного запроса будут обновлены все поля.
##### Request
Содержит структуру типа ```Feedback```
##### Response
В случае отсутствия фидбека будет возвращен соответствующий код GRPC

#### ListFeedbacksV1
Позволяет получить информацию о нескольких фидбеках.  
##### Request
```limit``` - максимальное количество возвращаемых фидбеков  
```offset``` - число строк, подлежащее пропуску.
##### Response
Содержит 1..N структур типа ```Feedback```  
В случае отсутствия фидбеков будет возвращен пустой ответ.

____











