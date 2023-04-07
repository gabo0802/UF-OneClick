import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';
import { SubscriptionListComponent } from './subscription-list/subscription-list.component';

import { UsersComponent } from './users.component';
import { WelcomeHeaderComponent } from './welcome-header/welcome-header.component';
import { ReportComponent } from './report/report.component';

describe('UsersComponent', () => {
  let component: UsersComponent;
  let fixture: ComponentFixture<UsersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        UsersComponent,
        SubscriptionListComponent,
        WelcomeHeaderComponent,
        ReportComponent
       ],
       imports: [
        HttpClientModule, 
        MaterialDesignModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UsersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
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
});
