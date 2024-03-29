export interface Subscription {
     name: string;
     price: number | string;
     dateadded: Date | string;
     dateremoved: Date | string;
     email: string;
     subid: string;
     timezone?: string;
     userid: number | string;
     username?: string;
     usersubid: string;     
}