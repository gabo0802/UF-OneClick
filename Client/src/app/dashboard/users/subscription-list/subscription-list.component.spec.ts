import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { SubscriptionListComponent } from './subscription-list.component';

describe('SubscriptionListComponent', () => {
  let component: SubscriptionListComponent;
  let fixture: ComponentFixture<SubscriptionListComponent>;  

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        SubscriptionListComponent
      ],
      imports: [
        MaterialDesignModule, 
        HttpClientModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
      ],
      providers: [        
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SubscriptionListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();    
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('subscription list should be initially empty', ()=> {

    expect(component.subscriptionList).toEqual([]);
  });

  it('subscription table should have initial column values', () => {

    expect(component.displayedColumns).toEqual(['name', 'price', 'dateadded', 'actions']);
  });

  it('getActive should change table display columns', () => {

    component.getActive();
    expect(component.displayedColumns).toEqual(['name', 'price', 'dateadded', 'actions']);
  });

  it('getActive should change active to true', () => {

    component.getActive();
    expect(component.active).toEqual(true);
  });

  it('getInactive should change table display columns', () => {

    component.getInactive();
    expect(component.displayedColumns).toEqual(['name', 'price', 'dateadded', 'dateremoved', 'actions']);
  });

  it('getInactive should change active to false', () => {

    component.getInactive();
    expect(component.active).toEqual(false);
  });
});
