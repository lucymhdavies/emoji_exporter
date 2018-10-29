.PHONY: plan apply clean

plan:
	@echo ========================================
	@echo "Getting dependencies"
	@echo ========================================
	lucli terraform get
	lucli terraform init
	@echo ========================================
	@echo "Pulling Remote Terraform State"
	@echo ========================================
	lucli terraform state pull
	@echo ========================================
	@echo Terraform Plan
	@echo ========================================
	lucli terraform plan -- -out plan.out
	@echo ========================================
	@echo Saving plaintext plan
	@echo ========================================
	lucli terraform show -- -no-color plan.out > plan.txt
	cat plan.txt | grep -e '^+\|^-\|^~' > plan.summary.txt || true

apply:
	@echo ========================================
	@echo Terraform Apply
	@echo ========================================
	lucli terraform apply plan.out

clean:
	rm -f plan.out plan.txt plan.summary.txt
