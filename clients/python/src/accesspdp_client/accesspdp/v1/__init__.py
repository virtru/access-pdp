# Hack because
# 1. Python still has a poopy module import system that doesn't rely on canonical identifiers
# and never quite works correctly in all cases
# 2. gRPC's python codegen sucks
import sys,os
sys.path.append(os.path.join(os.path.dirname(os.path.abspath(__file__)),os.pardir,os.pardir))
