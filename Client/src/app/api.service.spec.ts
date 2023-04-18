import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { ApiService } from './api.service';
import { HttpClientModule } from '@angular/common/http';
import { of } from 'rxjs';
import { Subscription } from './subscription.model';


describe('ApiService', () => {
  let service: ApiService;
  let httpTestingController: HttpTestingController;
  let httpClient: HttpClientModule;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        HttpClientTestingModule,        
      ]
    });

    service = TestBed.inject(ApiService);
    httpClient = TestBed.inject(HttpClientModule);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('can GET user information from getUserInfo', () => {
    
    const testUserData = { username: "bunny", timezone: "-400", email: "coolhacker@gmail.com"};

    service.getUserInfo().subscribe(data => {

      expect(data).toEqual(testUserData);
    })

    const req = httpTestingController.expectOne('api/alluserinfo');

    expect(req.request.method).toEqual('GET');

    req.flush(testUserData);

    httpTestingController.verify();
  });

  it('can GET user active Subscriptions from getActiveUserSubscriptions', () => {
    
    let subs: Subscription[] = [{ name: "Netflix", dateadded:"2022-10-05 11:10:45", dateremoved: "", price: "45.99", email: "", subid: "4", userid: "", usersubid: "45", username: ""}];
    const testSubsData = subs;

    service.getActiveUserSubscriptions().subscribe(data => {

      testSubsData[0].dateadded = new Date(testSubsData[0].dateadded);
      expect(data).toEqual(testSubsData);
    })

    const req = httpTestingController.expectOne('/api/subscriptions/active');

    expect(req.request.method).toEqual('GET');

    req.flush(testSubsData);

    httpTestingController.verify();
  });

  it('can GET user all Subscriptions (Services) from getAllSubscriptions', () => {
    
    let subs: Subscription[] = [
      { name: "Netflix", dateadded:"2022-10-05 11:10:45", dateremoved: "", price: "45.99", email: "", subid: "4", userid: "", usersubid: "45", username: ""},
      { name: "Hulu", dateadded:"2020-11-023 09:10:45", dateremoved: "", price: "9.99", email: "", subid: "5", userid: "", usersubid: "46", username: ""}
    ];
    const testSubsData = subs;

    service.getAllSubscriptions().subscribe(data => {

      expect(data).toEqual(testSubsData);
    })

    const req = httpTestingController.expectOne('/api/subscriptions/services');

    expect(req.request.method).toEqual('GET');

    req.flush(testSubsData);

    httpTestingController.verify();
  });

  it('can GET user all inactive Subscriptions from getAllInactiveUserSubscriptions', () => {
    
    let subs: Subscription[] = [
      { name: "Netflix", dateadded:"2020-10-05 11:10:45", dateremoved: "2022-10-05 11:10:49", price: "45.99", email: "", subid: "4", userid: "", usersubid: "45", username: ""},
      { name: "Hulu", dateadded:"2019-11-023 09:10:45", dateremoved: "2020-11-023 09:10:44", price: "9.99", email: "", subid: "5", userid: "", usersubid: "46", username: ""}
    ];
    const testSubsData = subs;

    service.getAllInactiveUserSubscriptions().subscribe(data => {

      testSubsData[0].dateadded = new Date(testSubsData[0].dateadded);
      testSubsData[1].dateadded = new Date(testSubsData[1].dateadded);
      testSubsData[0].dateremoved = new Date(testSubsData[0].dateremoved);
      testSubsData[1].dateremoved = new Date(testSubsData[1].dateremoved);

      expect(data).toEqual(testSubsData);
    })

    const req = httpTestingController.expectOne('/api/subscriptions');

    expect(req.request.method).toEqual('GET');

    req.flush(testSubsData);

    httpTestingController.verify();
  });

  it('can POST user information to login', () => {
    
    const testUserData = { username: "bunny", password: "secret"};

    service.login(testUserData).subscribe(data => {

      expect(data).toEqual(testUserData, 'should return user info')
    })

    const req = httpTestingController.expectOne('/api/login');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserData);

    httpTestingController.verify();
  });

  it('can POST user information to signup', () => {
    
    const testUserData = { username: "bunny", email: "fake@fake.com", password: "secret"};

    service.createUser(testUserData).subscribe(data => {

      expect(data).toEqual(testUserData, 'should return user info')
    })

    const req = httpTestingController.expectOne('/api/accountcreation');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserData);

    httpTestingController.verify();
  });

  it('can POST user subscription information to createUserSubscription', () => {
    
    const testUserSubData = { name: "AMC+", price: "14.99"};

    service.createUserSubscription(testUserSubData.name, testUserSubData.price).subscribe(data => {

      expect(data).toEqual(testUserSubData, 'should return user sub info')
    })

    const req = httpTestingController.expectOne('api/subscriptions/createsubscription');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can POST user subscription name to addActiveUserSubscription', () => {
    
    const testUserSubData = { name: "AMC+"};

    service.addActiveUserSubscription(testUserSubData.name).subscribe(data => {

      expect(data).toEqual(testUserSubData, 'should return user sub info')
    })

    const req = httpTestingController.expectOne('api/subscriptions/addsubscription');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can POST user subscription info to addOldUserSubscription', () => {
    
    const testUserSubData = { name: "AMC+", price: "14.99", dateadded: "2020-08-19 12:34:23", dateremoved:"2022-10-22 22:34:15"};

    service.addOldUserSubscription(testUserSubData).subscribe(data => {

      expect(data).toEqual(testUserSubData, 'should return user sub info')
    })

    const req = httpTestingController.expectOne('api/subscriptions/addoldsubscription');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can POST user subscription info to deactivateSubscription', () => {
    
    const testUserSubData = { name: "Paramount+"};

    service.deactivateSubscription(testUserSubData.name).subscribe(data => {

      expect(data).toEqual(testUserSubData)
    })

    const req = httpTestingController.expectOne('api/subscriptions/cancelsubscription');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can DELETE user subscription to deactivateSubscription', () => {
    
    const testUserSubData = { name: "Netflix", price: "14.99", dateadded: "2020-08-19 12:34:23", dateremoved:"2022-10-22 22:34:15", usersubid: "45"};

    service.deleteUserSubscription(testUserSubData.usersubid).subscribe(data => {

      expect(data).toEqual(testUserSubData.usersubid, 'should delete user sub')
    })

    const req = httpTestingController.expectOne(`api/subscriptions/${testUserSubData.usersubid}`);

    expect(req.request.method).toEqual('DELETE');

    req.flush(testUserSubData.usersubid);

    httpTestingController.verify();
  });

  it('can POST user subscription info to reactivateSubscription', () => {
    
    const testUserSubData = {name: "Paramount+"};

    service.reactivateSubscription(testUserSubData.name).subscribe(data => {

      expect(data).toEqual(testUserSubData)
    })

    const req = httpTestingController.expectOne('/api/subscriptions/addsubscription');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can PUT username info to updateUsername', () => {
    
    const testUserSubData = {username: "GatorMan"};

    service.updateUsername(testUserSubData.username).subscribe(data => {

      expect(data).toEqual(testUserSubData)
    })

    const req = httpTestingController.expectOne('api/changeusername');

    expect(req.request.method).toEqual('PUT');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can PUT user email info to updateUserEmail', () => {
    
    const testUserSubData = {email: "GatorMan@hotmail.com"};

    service.updateUserEmail(testUserSubData.email).subscribe(data => {

      expect(data).toEqual(testUserSubData)
    })

    const req = httpTestingController.expectOne('api/changeemail');

    expect(req.request.method).toEqual('PUT');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can PUT user passwords info to updateUserPassword', () => {
    
    const testUserSubData = {oldpassword: "abcd", newpassword: "1234"};

    service.updateUserPassword(testUserSubData.oldpassword, testUserSubData.newpassword).subscribe(data => {

      expect(data).toEqual(testUserSubData)
    })

    const req = httpTestingController.expectOne('api/changepassword');

    expect(req.request.method).toEqual('PUT');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can PUT user passwords info to updateTimezone', () => {
    
    const testUserSubData = {timezone: "0500"};

    service.updateTimezone(testUserSubData.timezone).subscribe(data => {

      expect(data).toEqual(testUserSubData)
    })

    const req = httpTestingController.expectOne('api/changetimezone');

    expect(req.request.method).toEqual('PUT');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can DELETE user account to deleteUserAccount', () => {
    
    const testUserSubData = {};

    service.deleteUserAccount().subscribe(data => {

      expect(data).toEqual({}, 'should delete user account')
    })

    const req = httpTestingController.expectOne('api/deleteuser');

    expect(req.request.method).toEqual('DELETE');

    req.flush(testUserSubData);

    httpTestingController.verify();
  });

  it('can POST dates info to graphPrices', () => {
    
    const testUserData = {startMonth: 3, startYear: 2019, endMonth: 4, endYear: 2021};

    service.graphPrices(testUserData.startMonth, testUserData.startYear, testUserData.endMonth, testUserData.endYear).subscribe(data => {

      expect(data).toEqual(testUserData)
    })

    const req = httpTestingController.expectOne('api/getallprices');

    expect(req.request.method).toEqual('POST');

    req.flush(testUserData);

    httpTestingController.verify();
  });
});
