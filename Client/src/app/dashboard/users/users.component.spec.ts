import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';
import { SubscriptionListComponent } from './subscription-list/subscription-list.component';
import { UsersComponent } from './users.component';
import { ReportComponent } from './report/report.component';
import { ApiService } from 'src/app/api.service';
import { of } from 'rxjs';
import { Subscription } from 'src/app/subscription.model';


describe('UsersComponent', () => {
  let component: UsersComponent;
  let fixture: ComponentFixture<UsersComponent>;
  let api: ApiService;  

  beforeEach(async () => {
    // const ApiServiceSpy = jasmine.createSpyObj<ApiService>(['getUserInfo', 'getActiveUserSubscriptions']);
    // ApiServiceSpy.getActiveUserSubscriptions.and.returnValue(of(new Array<Subscription>()));
    // ApiServiceSpy.getUserInfo.and.returnValue(of({username: ''}));    

    await TestBed.configureTestingModule({
      declarations: [ 
        UsersComponent,
        SubscriptionListComponent,
        ReportComponent
       ],
       imports: [
        HttpClientTestingModule, 
        MaterialDesignModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
       ],
       providers: [
        { provide: ApiService}
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UsersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    
    api = TestBed.inject(ApiService);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('username variable should initially be empty', () => {    
    expect(component.username).toEqual('');
  });

  it('subscriptionList should be initially be empty', ()=> {
    expect(component.subscriptionList.length).toEqual(0);
  });

  it('displaySubscriptions initial value should be true', ()=> {
    expect(component.displaySubscriptions).toEqual(true);
  });

  it('displayReport initial value should be false', ()=> {
    expect(component.displayReport).toEqual(false);
  });

  it('todayDate should be today', () => {

    let today: Date = new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate());

    expect(component.todayDate).toEqual(today);
  });

  it('user with username dragon should have username dragon', ()=> {

    spyOn(api, 'getUserInfo').and.returnValue(of({username: 'dragon'}))
    spyOn(api, 'getActiveUserSubscriptions').and.returnValue(of(new Array<Subscription>()))
  
    component.ngOnInit()
    expect(component.username).toEqual("dragon");
  });
  
  it('When subscriptions displayed, clicking Report side nav should set displaySubscriptions to false', () => {
    
    component.showReport();

    expect(component.displaySubscriptions).toEqual(false);
  });

  it('When subscriptions displayed, clicking Report side nav should set displayReport to true', () => {
    
    component.showReport();

    expect(component.displayReport).toEqual(true);
  });

  it('UpdateActiveSubscriptions should be called when switching to Subscriptions', ()=> {

    spyOn(component, 'updateActiveSubscriptions');
    component.showReport();
    component.showSubscriptions();
    expect(component.updateActiveSubscriptions).toHaveBeenCalled();
  });

});
