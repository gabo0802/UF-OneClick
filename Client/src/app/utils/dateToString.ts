export function dateToString(date: Date): string {
  let currentDate = new Date();

  let month: string = (date.getMonth() + 1).toString();
  let dayDate: string = date.getDate().toString();    
  let hours: string = currentDate.getHours().toString();
  let minutes: string = currentDate.getMinutes().toString();
  let seconds: string = currentDate.getSeconds().toString();

  if(currentDate.getHours() < 10){
    hours = '0' + hours
  }

  if(currentDate.getMinutes() < 10){
    minutes = '0' + minutes
  }

  if(currentDate.getSeconds() < 10){
      seconds = '0' + seconds
  }

  //if month is less than 10 adds leading zero
  if(date.getMonth() < 10){
    month = '0' + month;
  }

  //if date is less than 10 adds leading zero
  if(date.getDate() < 10) {
    
    dayDate = '0' + dayDate;
  }

  let stringDate = date.getFullYear() + "-" + month + "-" + dayDate + " " + hours + ":" + minutes + ":" + seconds;
  
  return stringDate;
}

export function dateToStringOffset(date: Date, secondsOffset: number): string {
  let currentDate = new Date();
  currentDate.setSeconds(currentDate.getSeconds() - secondsOffset)

  let month: string = (date.getMonth() + 1).toString();
  let dayDate: string = date.getDate().toString();

  let hours: string = currentDate.getHours().toString();
  let minutes: string = currentDate.getMinutes().toString();
  let seconds: string = (currentDate.getSeconds()).toString();

  if(currentDate.getHours() < 10){
    hours = '0' + hours
  }

  if(currentDate.getMinutes() < 10){
    minutes = '0' + minutes
  }

  if(currentDate.getSeconds() < 10){
      seconds = '0' + seconds
  }

  //if month is less than 10 adds leading zero
  if(date.getMonth() < 10){
    month = '0' + month;
  }

  //if date is less than 10 adds leading zero
  if(date.getDate() < 10) {
    
    dayDate = '0' + dayDate;
  }

  let stringDate = date.getFullYear() + "-" + month + "-" + dayDate + " " + hours + ":" + minutes + ":" + seconds;
  
  return stringDate;
}