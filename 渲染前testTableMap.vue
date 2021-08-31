<template>
  <div>
      databaseName : {[.appCode]}
  </div>
  <div>
    {[range .table]}
       <div>tableName :{[.name]}</div>
       {[range .field]}
            columnName: {[.name]}
            fieldType: {[.fieldType]}
       {[end]}
    {[end]}
  </div>
</template>