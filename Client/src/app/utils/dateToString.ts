export function dateToString(date: Date): string {

    //seconds between 10 and 60
    let randomSecs: number = Math.floor(Math.random() * 50 + 10);

    let month: string = (date.getMonth() + 1).toString();
    let dayDate: string = date.getDate().toString();    

    //if month is less than 10 adds leading zero
    if(date.getMonth() < 10){

      month = '0' + month;
    }

    //if date is less than 10 adds leading zero
    if(date.getDate() < 10) {
      
      dayDate = '0' + dayDate;
    }

    let stringDate = date.getFullYear() + "-" + month + "-" + dayDate + " 12:00:" + randomSecs;
    
    return stringDate;
}